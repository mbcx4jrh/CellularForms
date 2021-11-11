package main

type cell struct {
	index    int
	position Vector3
	links    []cell
}

type cellform struct {
	cells []cell
}

func createCellsStructure(maxCells int) *cellform {
	var cells []cell = make([]cell, maxCells)
	return &cellform{cells[0:0]}
}

func (c *cellform) seedMesh(m mesh) {

	seedCells := importMesh(m)
	c.cells = append(c.cells, seedCells...)
}
