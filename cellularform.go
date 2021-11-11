package main

type cell struct {
	position Vector3
}

type cellform struct {
	cells []cell
}

func createCellsStructure(maxCells int) *cellform {
	var cells []cell = make([]cell, maxCells)
	return &cellform{cells[0:0]}
}

func (c *cellform) seedMesh(m mesh) {
	for _, v := range m.vertices {
		c.cells = append(c.cells, cell{v})
	}
}
