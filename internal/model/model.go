package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int
	Content string
	Author  string
}

type Sheet struct {
	CalcChain                      []CalcChain                      `json:"calcChain,omitempty" bson:"calcChain,omitempty"`
	Celldata                       []Cell                           `json:"celldata,omitempty" bson:"celldata,omitempty"`
	ChWidth                        int                              `json:"ch_width,omitempty" bson:"ch_width,omitempty"`
	Column                         int                              `json:"column" bson:"column"`
	Config                         *SheetConfig                     `json:"config,omitempty" bson:"config,omitempty"`
	Index                          string                           `json:"index" bson:"index"`
	LuckysheetSelectSave           []SelectSave                     `json:"luckysheet_select_save,omitempty" bson:"luckysheet_select_save,omitempty"`
	Name                           string                           `json:"name" bson:"name"`
	Order                          int                              `json:"order" bson:"order"`
	RhHeight                       int                              `json:"rh_height,omitempty" bson:"rh_height,omitempty"`
	Row                            int                              `json:"row,omitempty" bson:"row,omitempty"`
	ScrollLeft                     int                              `json:"scrollLeft,omitempty" bson:"scrollLeft,omitempty"`
	ScrollTop                      int                              `json:"scrollTop,omitempty" bson:"scrollTop,omitempty"`
	Status                         FlexInt                          `json:"status" bson:"status"`
	ShowGridLines                  int                              `json:"showGridLines,omitempty" bson:"showGridLines,omitempty"`
	ZoomRatio                      float64                          `json:"zoomRatio,omitempty" bson:"zoomRatio,omitempty"`
	Filter                         map[string]Filter                `json:"filter,omitempty" bson:"filter,omitempty"`
	FilterSelect                   *FilterSelect                    `json:"filter_select,omitempty" bson:"filter_select,omitempty"`
	Images                         interface{}                      `json:"images,omitempty" bson:"images,omitempty"`
	AlternateFormatSave            []AlternateFormatSave            `json:"luckysheet_alternateformat_save,omitempty" bson:"luckysheet_alternateformat_save,omitempty"`
	AlternateFormatSaveModelCustom []AlternateFormatSaveModelCustom `json:"luckysheet_alternateformat_save_modelCustom,omitempty" bson:"luckysheet_alternateformat_save_modelCustom,omitempty"`
	ConditionFormatSave            []ConditionFormatSave            `json:"luckysheet_conditionformat_save,omitempty" bson:"luckysheet_conditionformat_save,omitempty"`
	Frozen                         *Frozen                          `json:"frozen,omitempty" bson:"frozen,omitempty"`
	Chart                          []Chart                          `json:"chart,omitempty" bson:"chart,omitempty"`
	Image                          []Image                          `json:"image,omitempty" bson:"image,omitempty"`
	IsPivotTable                   bool                             `json:"isPivotTable,omitempty" bson:"isPivotTable,omitempty"`
	PivotTable                     *PivotTable                      `json:"pivotTable,omitempty" bson:"pivotTable,omitempty"`
	DynamicArray                   []DynamicArray                   `json:"dynamicArray,omitempty" bson:"dynamicArray,omitempty"`
}
