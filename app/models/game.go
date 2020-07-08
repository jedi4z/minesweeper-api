package models

type Game struct {
	CommonFields
	NumberOfCols  int    `json:"number_of_cols" sql:"type:int"`
	NumberOfRows  int    `json:"number_of_rows" sql:"type:int"`
	NumberOfMines int    `json:"number_of_mines" sql:"type:int"`
	Grid          []*Row `json:"grid,omitempty" sql:"foreignkey:GameID"`
}
