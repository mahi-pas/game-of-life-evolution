package models

type GameBoard struct {
	Seed        string
	BoardWidth  int
	BoardHeight int
	Board       [][]bool
	Generations int
}

func generateBoard(Seed string, BoardWidth int, BoardHeight int)
