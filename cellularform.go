package main

type cell struct {
	index    int
	position Vector3
	links    []cell
}

type cellformParams struct {
	linkLength      float64
	springFactor    float64
	planarFactor    float64
	bulgeFactor     float64
	repulsionRange  float64
	repulsionFactor float64
}

type cellform struct {
	cells []cell
}

func NewCellform(maxCells int) *cellform {
	var cells []cell = make([]cell, maxCells)
	return &cellform{cells[0:0]}
}

func (c *cellform) seedMesh(m mesh) {
	seedCells := importMesh(m)
	c.cells = append(c.cells, seedCells...)
}

func (c *cellform) iterate() {

}
