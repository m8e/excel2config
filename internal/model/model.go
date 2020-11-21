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

type MergeCell struct {
	C  int `json:"c,omitempty" bson:"c,omitempty"`
	Cs int `json:"cs,omitempty" bson:"cs,omitempty"`
	R  int `json:"r,omitempty" bson:"r,omitempty"`
	Rs int `json:"rs,omitempty" bson:"rs,omitempty"`
}

type PS struct {
	Height int    `json:"height,omitempty" bson:"height,omitempty"`
	Width  int    `json:"width,omitempty" bson:"width,omitempty"`
	Left   int    `json:"left,omitempty" bson:"left,omitempty"`
	Top    int    `json:"top,omitempty" bson:"top,omitempty"`
	IsShow bool   `json:"isshow,omitempty" bson:"isshow,omitempty"`
	Value  string `json:"value,omitempty" bson:"value,omitempty"`
}

type CellType struct { // celltype 单元格值格式
	Fa string `json:"fa,omitempty" bson:"fa,omitempty"`
	S  []struct {
		Bl int         `json:"bl,omitempty" bson:"bl,omitempty"`
		Cl int         `json:"cl,omitempty" bson:"cl,omitempty"`
		Fc string      `json:"fc,omitempty" bson:"fc,omitempty"`
		Ff string      `json:"ff,omitempty" bson:"ff,omitempty"`
		Fs interface{} `json:"fs,omitempty" bson:"fs,omitempty"`
		It int         `json:"it,omitempty" bson:"it,omitempty"`
		Un int         `json:"un,omitempty" bson:"un,omitempty"`
		V  string      `json:"v,omitempty" bson:"v,omitempty"`
	} `json:"s,omitempty" bson:"s,omitempty"`
	T string `json:"t,omitempty" bson:"t,omitempty"`
}

type CellValue struct {
	Bg string      `json:"bg,omitempty" bson:"bg,omitempty"` // 背景颜色
	Bl int         `json:"bl,omitempty" bson:"bl,omitempty"` // 0 常规 、 1加粗
	Cl int         `json:"cl,omitempty" bson:"cl,omitempty"` // 删除线
	Ct *CellType   `json:"ct,omitempty" bson:"ct,omitempty"`
	F  string      `json:"f,omitempty" bson:"f,omitempty"`   // 公式
	Fc string      `json:"fc,omitempty" bson:"fc,omitempty"` // 字体颜色
	Ff int         `json:"ff,omitempty" bson:"ff,omitempty"` // 字体类型
	Fs interface{} `json:"fs,omitempty" bson:"fs,omitempty"` // 字体大小
	It int         `json:"it,omitempty" bson:"it,omitempty"` // 斜体
	M  string      `json:"m,omitempty" bson:"m,omitempty"`   // 显示值
	Mc *MergeCell  `json:"mc,omitempty" bson:"mc,omitempty"` // 合并单元格
	Tb string      `json:"tb,omitempty" bson:"tb,omitempty"` // 文本换行，0 截断、1溢出、2 自动换行
	Tr interface{} `json:"tr,omitempty" bson:"tr,omitempty"` // 竖排文字
	V  interface{} `json:"v,omitempty" bson:"v,omitempty"`   // 原始值
	Ht interface{} `json:"ht,omitempty" bson:"ht,omitempty"` // 水平对齐，0 居中、1 左、2右
	Vt interface{} `json:"vt,omitempty" bson:"vt,omitempty"` // 垂直对齐，0 中间、1 上、2下
	Rt interface{} `json:"rt,omitempty" bson:"rt,omitempty"` // 文字旋转角度
	Ps *PS         `json:"ps,omitempty" bson:"ps,omitempty"` //批注
}

type Cell struct {
	C int       `json:"c" bson:"c"`
	R int       `json:"r" bson:"r"`
	V CellValue `json:"v,omitempty" bson:"v,omitempty"`
}

type BorderValueStyle struct {
	Color string `json:"color,omitempty" bson:"color,omitempty"`
	Style int    `json:"style,omitempty" bson:"style,omitempty"`
}

type Border struct {
	BorderType string `json:"borderType,omitempty" bson:"borderType,omitempty"`
	Color      string `json:"color,omitempty" bson:"color,omitempty"`
	Range      []struct {
		Column []int `json:"column,omitempty" bson:"column,omitempty"`
		Row    []int `json:"row,omitempty" bson:"row,omitempty"`
	} `json:"range,omitempty" bson:"range,omitempty"`
	RangeType string `json:"rangeType,omitempty" bson:"rangeType,omitempty"`
	Style     string `json:"style,omitempty" bson:"style,omitempty"`
	Value     struct {
		RowIndex int              `json:"row_index,omitempty" bson:"row_index,omitempty"`
		ColIndex int              `json:"col_index,omitempty" bson:"col_index,omitempty"`
		B        BorderValueStyle `json:"b,omitempty" bson:"b,omitempty"`
		L        BorderValueStyle `json:"l,omitempty" bson:"l,omitempty"`
		R        BorderValueStyle `json:"r,omitempty" bson:"r,omitempty"`
		T        BorderValueStyle `json:"t,omitempty" bson:"t,omitempty"`
	} `json:"value,omitempty" bson:"value,omitempty"`
}

type ConfigMerge struct {
	C  int `json:"c,omitempty" bson:"c,omitempty"`
	Cs int `json:"cs,omitempty" bson:"cs,omitempty"`
	R  int `json:"r,omitempty" bson:"r,omitempty"`
	Rs int `json:"rs,omitempty" bson:"rs,omitempty"`
}

type SheetConfig struct {
	BorderInfo      []Border               `json:"borderInfo,omitempty" bson:"borderInfo,omitempty"`
	Columnlen       map[string]int         `json:"columnlen,omitempty" bson:"columnlen,omitempty"`
	CurentsheetView string                 `json:"curentsheetView,omitempty" bson:"curentsheetView,omitempty"`
	CustomHeight    map[string]int         `json:"customHeight,omitempty" bson:"customHeight,omitempty"`
	CustomWidth     map[string]int         `json:"customWidth,omitempty" bson:"customWidth,omitempty"`
	Merge           map[string]ConfigMerge `json:"merge,omitempty" bson:"merge,omitempty"`
	Rowhidden       map[string]int         `json:"rowhidden,omitempty" bson:"rowhidden,omitempty"`
	Colhidden       map[string]int         `json:"colhidden,omitempty" bson:"colhidden,omitempty"`
	Rowlen          map[string]int         `json:"rowlen,omitempty" bson:"rowlen,omitempty"`
	SheetViewZoom   struct {
		ViewLayoutZoomScale int     `json:"viewLayoutZoomScale,omitempty" bson:"viewLayoutZoomScale,omitempty"`
		ViewNormalZoomScale int     `json:"viewNormalZoomScale,omitempty" bson:"viewNormalZoomScale,omitempty"`
		ViewPageZoomScale   float64 `json:"viewPageZoomScale,omitempty" bson:"viewPageZoomScale,omitempty"`
	} `json:"sheetViewZoom,omitempty" bson:"sheetViewZoom,omitempty"`
	Authority interface{} `json:"authority,omitempty" bson:"authority,omitempty"`
}

type Sheet struct {
	CalcChain []struct {
		C       int           `json:"c,omitempty" bson:"c,omitempty"`
		Chidren struct{}      `json:"chidren,omitempty" bson:"chidren,omitempty"`
		Color   string        `json:"color,omitempty" bson:"color,omitempty"`
		Func    []interface{} `json:"func,omitempty" bson:"func,omitempty"`
		Index   string        `json:"index,omitempty" bson:"index,omitempty"`
		Parent  interface{}   `json:"parent,omitempty" bson:"parent,omitempty"`
		R       int           `json:"r,omitempty" bson:"r,omitempty"`
		Times   int           `json:"times,omitempty" bson:"times,omitempty"`
	} `json:"calcChain,omitempty" bson:"calcChain,omitempty"`
	Celldata             []Cell      `json:"celldata,omitempty" bson:"celldata,omitempty"`
	ChWidth              int         `json:"ch_width,omitempty" bson:"ch_width,omitempty"`
	Column               int         `json:"column,omitempty" bson:"column,omitempty"`
	Config               SheetConfig `json:"config,omitempty" bson:"config,omitempty"`
	Index                string      `json:"index,omitempty" bson:"index,omitempty"`
	LuckysheetSelectSave []struct {
		Column      []int `json:"column,omitempty" bson:"column,omitempty"`
		ColumnFocus int   `json:"column_focus,omitempty" bson:"column_focus,omitempty"`
		Height      int   `json:"height,omitempty" bson:"height,omitempty"`
		HeightMove  int   `json:"height_move,omitempty" bson:"height_move,omitempty"`
		Left        int   `json:"left,omitempty" bson:"left,omitempty"`
		LeftMove    int   `json:"left_move,omitempty" bson:"left_move,omitempty"`
		Row         []int `json:"row,omitempty" bson:"row,omitempty"`
		RowFocus    int   `json:"row_focus,omitempty" bson:"row_focus,omitempty"`
		Top         int   `json:"top,omitempty" bson:"top,omitempty"`
		TopMove     int   `json:"top_move,omitempty" bson:"top_move,omitempty"`
		Width       int   `json:"width,omitempty" bson:"width,omitempty"`
		WidthMove   int   `json:"width_move,omitempty" bson:"width_move,omitempty"`
	} `json:"luckysheet_select_save,omitempty" bson:"luckysheet_select_save,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Order      string `json:"order,omitempty" bson:"order,omitempty"`
	RhHeight   int    `json:"rh_height,omitempty" bson:"rh_height,omitempty"`
	Row        int    `json:"row,omitempty" bson:"row,omitempty"`
	ScrollLeft int    `json:"scrollLeft,omitempty" bson:"scrollLeft,omitempty"`
	ScrollTop  int    `json:"scrollTop,omitempty" bson:"scrollTop,omitempty"`
	Status     int    `json:"status,omitempty" bson:"status,omitempty"`
	ZoomRatio  int    `json:"zoomRatio,omitempty" bson:"zoomRatio,omitempty"`
}
