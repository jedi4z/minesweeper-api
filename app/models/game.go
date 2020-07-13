package models

import (
	"math/rand"
)

type Game struct {
	CommonFields
	Status        string `json:"status" sql:"type:varchar(10)"`
	NumberOfCols  int    `json:"number_of_cols" sql:"type:int;not null" binding:"required"`
	NumberOfRows  int    `json:"number_of_rows" sql:"type:int;not null" binding:"required"`
	NumberOfMines int    `json:"number_of_mines" sql:"type:int;not null" binding:"required"`
	Grid          []*Row `json:"grid,omitempty" sql:"foreignkey:GameID"`
	User          *User  `json:"-"`
	UserID        uint   `json:"user_id"`
}

// Changes the Game status to OnHold
func (g *Game) HoldGame() {
	g.Status = OnHoldState
}

// Changes the Game status to Playing
func (g *Game) ResumeGame() {
	g.Status = PlayingState
}

// Creates a new grid for a Game
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

// Seeds the mines in the grid
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

// Counts how many neighbors there are near
// to all cells in the grid
func (g *Game) CountNeighbors() {
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]
			cell.CountMinesAround(g)
		}
	}
}

// Initialize the Game creating a grid, seeding the mines
// and counting the mines near to the cells
func (g *Game) InitGame() {
	g.CreateGrid()
	g.SeedMines()
	g.CountNeighbors()

	g.Status = PlayingState
}

func (g *Game) FlagCell(cellID uint) error {
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]

			if cell.ID == cellID {
				cell.Fag()
				return nil
			}
		}
	}

	return errCellNotFound
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

	return errCellNotFound
}

func (g *Game) CheckIfWon() {
	uncoveredCells := 0
	for i := 0; i < g.NumberOfRows; i++ {
		for j := 0; j < g.NumberOfCols; j++ {
			cell := g.Grid[i].Cells[j]

			if cell.Status == UncoveredState {
				uncoveredCells++
			}
		}
	}

	// if all non-mine cells are uncovered the game status is changed to won
	if uncoveredCells == ((g.NumberOfRows * g.NumberOfCols) - g.NumberOfMines) {
		g.Status = WonState
	}
}
