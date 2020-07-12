package models

type Cell struct {
	CommonFields
	HasMine     bool   `json:"has_mine" sql:"type:bool"`
	Status      string `json:"status" sql:"type:varchar(10)"`
	MinesAround int    `json:"mines_around" sql:"type:int"`
	RowIndex    int    `json:"row_index" sql:"type:int"`
	ColIndex    int    `json:"col_index" sql:"type:int"`
	RowID       uint   `json:"-"`
}

func isOnMap(row, col int, game *Game) bool {
	return row >= 0 && col >= 0 && row < game.NumberOfRows && col < game.NumberOfCols
}

func (c *Cell) Uncover() {
	c.Status = UncoveredState
	//
	//if c.MinesAround == 0 && c.Status == CoveredState {
	//	c.UncoverNeighbors(game)
	//}
}

func (c *Cell) CountMinesAround(game *Game) {
	mines := 0

	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			i, j := c.RowIndex+xOff, c.ColIndex+yOff

			if isOnMap(i, j, game) {
				cell := game.Grid[i].Cells[j]

				if cell.HasMine {
					mines++
				}
			}
		}
	}

	c.MinesAround = mines
}

func (c *Cell) UncoverNeighbors(game *Game) {
	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			i, j := c.RowIndex+xOff, c.ColIndex+yOff

			if isOnMap(i, j, game) {
				cell := game.Grid[i].Cells[j]

				if !cell.HasMine && cell.Status == CoveredState {
					cell.Status = UncoveredState
				}
			}
		}
	}
}
