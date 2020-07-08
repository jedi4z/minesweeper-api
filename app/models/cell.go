package models

type Cell struct {
	CommonFields
	HasMine     bool `json:"has_mine"`
	MinesAround int  `json:"mines_around,omitempty"`
	Row         int  `json:"row"`
	Col         int  `json:"col"`
	RowID       uint `json:"-"`
}
