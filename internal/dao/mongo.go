package dao

import (
	"context"
	"excel2config/internal/model"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const dbname = "sheets"

func NewMongo() (m *mongo.Client, cf func(), err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://fandy:fandypeng@127.0.0.1:27017/?authSource=sheets"))
	if err != nil {
		return nil, func() {}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.With("err", err).Error("mongo db ping failed")
		return
	}
	cf = func() { client.Disconnect(context.Background()) }
	m = client
	return
}

func (d *dao) PingMongo(ctx context.Context) error {
	if err := d.mongo.Ping(ctx, readpref.Primary()); err != nil {
		log.With("err", err).Error("mongo db ping failed")
		return err
	}
	return nil
}

func (d *dao) LoadExcel(ctx context.Context, gridKey string) (sheets []*model.Sheet, err error) {
	c := d.mongo.Database(dbname).Collection(gridKey)
	//先查出status=1的sheet的完整内容，其他status=0的sheet只取出基础信息，不包含sheet内容
	singleRes := c.FindOne(ctx, bson.D{{"status", 1}})
	activeSheet := new(model.Sheet)
	err = singleRes.Decode(activeSheet)
	if err != nil {
		log.With("err", err).Errorln("mongo decode error")
		return nil, err
	}
	opt := options.Find()
	opt.SetProjection(bson.D{{"name", true}, {"index", true}, {"order", true}, {"status", true}})
	opt.SetSort(bson.D{{"order", 1}})
	corsor, err := c.Find(ctx, bson.D{{"status", 0}}, opt)
	if err != nil {
		return
	}
	sheets = make([]*model.Sheet, 0)
	err = corsor.All(ctx, &sheets)
	if err != nil {
		log.With("err", err).Errorln("mango decode error")
		return
	}
	sheets = append(sheets, activeSheet)
	return
}

func (d *dao) LoadExcelSheet(ctx context.Context, gridKey string, indexs []string) (sheets map[string][]model.Cell, err error) {
	c := d.mongo.Database(dbname).Collection(gridKey)
	filters := make([]bson.M, 0)
	for _, index := range indexs {
		filters = append(filters, bson.M{"index": index})
	}
	opt := options.Find()
	opt.SetProjection(bson.D{{"celldata", true}, {"index", true}})
	corsor, err := c.Find(ctx, bson.M{"$or": filters}, opt)
	if err != nil {
		return
	}
	sheetInfos := make([]*model.Sheet, 0)
	err = corsor.All(ctx, &sheetInfos)
	if err != nil {
		log.With("err", err).Errorln("mango decode error")
		return
	}
	sheets = make(map[string][]model.Cell)
	for _, sheet := range sheetInfos {
		sheets[sheet.Index] = sheet.Celldata
	}
	return
}

func (d *dao) UpdateGridValue(ctx context.Context, gridKey string, req *model.UpdateV) (err error) {
	c := d.mongo.Database(dbname).Collection(gridKey)
	filter := bson.M{"index": req.I, "celldata.r": req.R, "celldata.c": req.C}
	formatV, err := d.format2Bson(req.V)
	if err != nil {
		log.With("err", err).With("v", req.V).Errorln("format2bson error")
		return
	}
	res, err := c.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"celldata.$.v", formatV}}}})
	if err != nil {
		log.With("err", err).With("gridKey", gridKey).With("req", req).Errorln("update error")
		return
	}
	if res.ModifiedCount <= 0 {
		formatCell, err := d.format2Bson(req.Cell)
		if err != nil {
			log.With("err", err).With("cell", req.Cell).Errorln("format2bson error")
			return err
		}
		_, err = c.UpdateOne(ctx, bson.M{"index": req.I}, bson.M{"$push": bson.M{"celldata": formatCell}})
	}
	return err
}

func (d *dao) UpdateGridMulti(ctx context.Context, gridKey string, req *model.UpdateRV) (err error) {
	subIndex := 0
	cs, ce := req.Range.Column[0], req.Range.Column[1]
	rs, re := req.Range.Row[0], req.Range.Row[1]
	for c := cs; c <= ce; c++ {
		index := 0
		for r := rs; r <= re; r++ {
			customReq := &model.UpdateV{
				Cell: model.Cell{
					C: c,
					R: r,
					V: req.V[index][subIndex],
				},
				I: req.I,
			}
			log.With("c", c+1).With("r", r+1).With("v", req.V[index][subIndex]).Infoln("update grid")
			err = d.UpdateGridValue(ctx, gridKey, customReq)
			if err != nil {
				log.With("err", err).With("req", customReq).Errorln("update error")
				return
			}
			index++
		}
		subIndex++
	}
	return err
}

func (d *dao) UpdateGridConfig(ctx context.Context, gridKey string, req *model.UpdateCG) (err error) {
	c := d.mongo.Database(dbname).Collection(gridKey)
	filter := bson.M{"index": req.I}
	formatV, err := d.format2Bson(map[string]interface{}{
		"config." + req.K: req.V,
	})
	if err != nil {
		log.With("err", err).With("v", req.V).Errorln("format2bson error")
		return
	}
	_, err = c.UpdateOne(ctx, filter, bson.D{{"$set", formatV}})
	if err != nil {
		log.With("err", err).With("gridKey", gridKey).With("req", req).Errorln("update error")
		return
	}
	return err
}

func (d *dao) UpdateGridCommon(ctx context.Context, gridKey string, req *model.UpdateCommon) (err error) {
	c := d.mongo.Database(dbname).Collection(gridKey)
	filter := bson.M{"index": req.I}
	formatV, err := d.format2Bson(map[string]interface{}{
		req.K: req.V,
	})
	if err != nil {
		log.With("err", err).With("v", req.V).Errorln("format2bson error")
		return
	}
	_, err = c.UpdateOne(ctx, filter, bson.D{{"$set", formatV}})
	if err != nil {
		log.With("err", err).With("gridKey", gridKey).With("req", req).Errorln("update error")
		return
	}
	return err
}

func (d *dao) format2Bson(v interface{}) (doc bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
