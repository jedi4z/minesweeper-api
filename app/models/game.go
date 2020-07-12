package models

import (
	"math/rand"
)

type Game struct {
	CommonFields
	Status        string `json:"status" sql:"type:varchar(10)"`
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
				Status:   CoveredState,
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

func (g *Game) CountNeighbors() {
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]
			cell.CountMinesAround(g)
		}
	}
}

func (g *Game) InitGame() {
	g.CreateGrid()
	g.SeedMines()
	g.CountNeighbors()
}

func (g *Game) GameOver() {
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]
			cell.Uncover()
		}
	}

	g.Status = GameOverState
}

func (g *Game) UncoverCell(cellID uint) error {
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]

			if cell.ID == cellID {
				// if the cell uncover has a mine the game is over
				if cell.HasMine {
					g.GameOver()
					return nil
				}

				// uncover the cell selected
				cell.Uncover()

				// uncover the neighbor's cells if
				// it doesn't have adjacent cells with mines
				if cell.MinesAround == 0 {
					cell.UncoverNeighbors(g)
				}

				return nil
			}
		}
	}

	return cellNotFound
}
