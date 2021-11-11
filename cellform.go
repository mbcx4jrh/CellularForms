package main

type cell struct {
	position Vector3
}

func createCellsStructure(maxCells int) []cell {
	cells := make([]cell, maxCells)
	return cells[0:1]
}
