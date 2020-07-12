package models

import "math/rand"

type Game struct {
	CommonFields
	NumberOfCols  int    `json:"number_of_cols" sql:"type:int"`
	NumberOfRows  int    `json:"number_of_rows" sql:"type:int"`
	NumberOfMines int    `json:"number_of_mines" sql:"type:int"`
	Grid          []*Row `json:"grid,omitempty" sql:"foreignkey:GameID"`
}

func (g *Game) CreateGrid() {
	rows := make([]*Row, 0)
	for rowIndex := 0; rowIndex < g.NumberOfRows; rowIndex++ {

		cells := make([]*Cell, 0)
		for colIndex := 0; colIndex < g.NumberOfCols; colIndex++ {
			cells = append(cells, &Cell{
				Status:   UncoveredState,
				RowIndex: rowIndex,
				ColIndex: colIndex,
			})
		}

		rows = append(rows, &Row{Cells: cells})
	}

	g.Grid = rows
}

func (g *Game) SeedMines() {
	k := 0
	for k < g.NumberOfMines {
		row := rand.Intn(g.NumberOfRows)
		col := rand.Intn(g.NumberOfCols)
		cell := g.Grid[row].Cells[col]

		if !cell.HasMine {
			cell.HasMine = true
			k++
		}
	}
}
