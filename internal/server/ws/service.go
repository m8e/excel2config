package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"excel2config/internal/dao"
	"excel2config/internal/model"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
	"strconv"
	"sync"
	"time"
)

type handler func(ctx context.Context, reqmsg []byte)

type service struct {
	*Client
	d        dao.Dao
	handlers map[string]handler

	sync.RWMutex
}

func newService(c *Client, d dao.Dao) *service {
	s := &service{
		Client: c,
		d:      d,
	}
	s.handlers = initHandlers(s)
	return s
}

func initHandlers(s *service) map[string]handler {
	return map[string]handler{
		"v":   s.updateGrid,
		"rv":  s.updateGridMulti,
		"cg":  s.updateGridConfig,
		"all": s.updateGridCommon,
	}
}

func (s *service) readAndServe() {
	log.Debugln("new ws connect")
	defer func() {
		s.Close()
	}()
	s.setReadOpts()
	for {
		messageType, message, err := s.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorln("ws socket error: %v", err)
			}
			break
		}
		log.Infoln("uid: ", s.uid, " recv message_type: ", messageType, ", msg: ", string(message))
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		s.handleRequest(context.Background(), message)
	}
}

func (s *service) handleRequest(ctx context.Context, reqmsg []byte) {
	if s.isClosed {
		return
	}
	var msg struct {
		T string `json:"t"`
	}
	var rsp struct {
		Type     int    `json:"type"`
		Id       int    `json:"id,omitempty"`
		UserName string `json:"username,omitempty"`
		Data     string `json:"data"`
	}
	json.Unmarshal(reqmsg, &msg)
	uid := s.GetUid()
	switch msg.T {
	case "v", "rv", "cg", "all", "fc", "drc", "arc", "f", "fsc", "fsr", "sha", "shc", "shd", "shr", "shre", "sh", "c", "na":
		rsp.Type = 2
	case "mv":
		rsp.Type = 3
		rsp.Id = uid
		rsp.UserName = "Guest" + strconv.Itoa(uid)
	case "rv_end": //离线情况下把更新指令打包批量下发给客户端
		rsp.Type = 4
	default:
		rsp.Type = 1
	}

	s.RLock()
	handler, ok := s.handlers[msg.T]
	s.RUnlock()
	if ok {
		handler(ctx, reqmsg)
	}

	rsp.Data = string(reqmsg)
	jsonstr, _ := json.Marshal(rsp)
	mgr.Send2AllClients(s.Client, jsonstr)
}

func (s *service) updateGrid(ctx context.Context, reqmsg []byte) {
	req := new(model.UpdateV)
	err := json.Unmarshal(reqmsg, req)
	if err != nil {
		log.With("err", err).With("jsonstr", string(reqmsg)).Errorln("json unmarshal error")
		return
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err = s.d.UpdateGridValue(ctx, s.gridKey, req)
	if err != nil {
		log.With("err", err).With("gridKey", s.gridKey).With("req", req).Errorln("update grid failed")
	}
}

func (s *service) updateGridMulti(ctx context.Context, reqmsg []byte) {
	req := new(model.UpdateRV)
	err := json.Unmarshal(reqmsg, req)
	if err != nil {
		log.With("err", err).With("jsonstr", string(reqmsg)).Errorln("json unmarshal error")
		return
	}
	if len(req.Range.Column) < 2 || len(req.Range.Row) < 2 {
		log.With("req", req).Warnln("invalid params")
		return
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	err = s.d.UpdateGridMulti(ctx, s.gridKey, req)
	if err != nil {
		log.With("err", err).With("gridKey", s.gridKey).With("req", req).Errorln("update grid failed")
	}
}

func (s *service) updateGridConfig(ctx context.Context, reqmsg []byte) {
	req := new(model.UpdateCG)
	err := json.Unmarshal(reqmsg, req)
	if err != nil {
		log.With("err", err).With("jsonstr", string(reqmsg)).Errorln("json unmarshal error")
		return
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	err = s.d.UpdateGridConfig(ctx, s.gridKey, req)
	if err != nil {
		log.With("err", err).With("gridKey", s.gridKey).With("req", req).Errorln("update grid failed")
	}
}

func (s *service) updateGridCommon(ctx context.Context, reqmsg []byte) {
	req := new(model.UpdateCommon)
	err := json.Unmarshal(reqmsg, req)
	if err != nil {
		log.With("err", err).With("jsonstr", string(reqmsg)).Errorln("json unmarshal error")
		return
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	err = s.d.UpdateGridCommon(ctx, s.gridKey, req)
	if err != nil {
		log.With("err", err).With("gridKey", s.gridKey).With("req", req).Errorln("update grid failed")
	}
}
