package model

type UpdateV struct {
	Cell
	I string `json:"i"`
	T string `json:"t"`
}

type UpdateRV struct {
	I     string `json:"i"`
	Range struct {
		Column []int `json:"column"`
		Row    []int `json:"row"`
	} `json:"range"`
	T string        `json:"t"`
	V [][]CellValue `json:"v"`
}

type UpdateCG struct {
	T string      `json:"t"`
	I string      `json:"i"`
	K string      `json:"k"`
	V interface{} `json:"v"`
}

type UpdateCommon struct {
	T string      `json:"t"`
	I string      `json:"i"`
	K string      `json:"k"`
	V interface{} `json:"v"`
}
