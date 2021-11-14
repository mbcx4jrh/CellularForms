package main

import "fmt"

type cellformParams struct {
	linkLength      float64
	springFactor    float64
	planarFactor    float64
	bulgeFactor     float64
	repulsionRange  float64
	repulsionFactor float64
}

type cellform struct {
	cells  []cell
	params cellformParams
}

func NewCellform(maxCells int, params cellformParams) *cellform {
	var cells []cell = make([]cell, maxCells)
	return &cellform{cells[0:0], params}
}

func (c *cellform) seedMesh(m mesh) {
	seedCells := importMesh(m)
	c.cells = append(c.cells, seedCells...)
}

func (c *cellform) iterate() {
	for _, cell := range c.cells {
		cell.computeNormal()
	}
}

func (c cellformParams) asString() string {
	return fmt.Sprintf("linkLength = %f\n"+
		"springFactor = %f\n"+
		"planarfactor = %f\n"+
		"bulgeFactor = %f\n"+
		"repulsionRange = %f\n"+
		"repulsionFactor = %f",
		c.linkLength, c.springFactor, c.planarFactor, c.bulgeFactor, c.repulsionRange, c.repulsionFactor)
}
