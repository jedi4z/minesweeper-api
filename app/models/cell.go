package models

type Cell struct {
	CommonFields
	HasMine     bool `json:"has_mine"`
	MinesAround int  `json:"mines_around,omitempty"`
	RowIndex    int  `json:"row"`
	ColIndex    int  `json:"col"`
	RowID       uint `json:"-"`
}

func isOnMap(row, col int, game *Game) bool {
	return row >= 0 && col >= 0 && row < game.NumberOfRows && col < game.NumberOfCols
}

func (c *Cell) CountMinesAround(game *Game) {
	mines := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if isOnMap(c.RowIndex+i, c.ColIndex+j, game) {
				cell := game.Grid[c.RowIndex+i].Cells[c.ColIndex+j]

				if cell.HasMine {
					mines++
				}
			}
		}
	}

	c.MinesAround = mines
}
