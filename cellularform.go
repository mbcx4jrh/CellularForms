package main

import (
	"fmt"
	"math"
	"time"

	"github.com/mbcx4jrh/vec3"
)

type cellformParams struct {
	linkLength      float64
	springFactor    float64
	planarFactor    float64
	bulgeFactor     float64
	repulsionRange  float64
	repulsionFactor float64
	feedRate        float64
}

type cellform struct {
	cells    []cell
	params   cellformParams
	maxCells int
}

func NewCellform(maxCells int, params cellformParams) *cellform {
	var cells []cell = make([]cell, maxCells)
	return &cellform{cells[0:0], params, maxCells}
}

func (c *cellform) seedMesh(m mesh) {
	seedCells := importMesh(m)
	c.cells = append(c.cells, seedCells...)
}

func (c *cellform) iterate() {
	c.checkForSplits()
	debug(fmt.Sprintf("Iterating through %d cells", len(c.cells)))
	link2 := c.params.linkLength * c.params.linkLength
	r2 := c.params.repulsionRange * c.params.repulsionRange
	start := time.Now()
	//for _, cell := range c.cells {
	for i := 0; i < len(c.cells); i++ {
		cell := &c.cells[i]
		cell.computeNormal()
		d_spring_sum := vec3.Zero()
		d_planar_sum := vec3.Zero()
		d_bulge_sum := 0.0
		d_collision_sum := vec3.Zero()

		for _, n := range cell.links {
			vectorBetween := vec3.Subtract(n.position, cell.position)
			vbNormalised := vec3.Normalize(vectorBetween)
			distance2 := vectorBetween.LengthSqr()
			d_spring_sum = vec3.Add(d_spring_sum,
				vec3.Subtract(vectorBetween,
					vec3.Mult(vbNormalised, c.params.linkLength)))
			d_planar_sum = vec3.Add(d_planar_sum, vectorBetween)
			if distance2 < link2 {
				dot := vec3.Dot(vectorBetween, cell.normal)
				d_bulge_sum += math.Sqrt(link2-distance2+dot*dot) + dot
			}
			d_collision_sum = vec3.Add(d_collision_sum, vec3.Mult(vbNormalised, (r2-distance2)/r2))
		}
		n := float64(len(cell.links))
		d_spring := vec3.Div(d_spring_sum, n)
		d_planar := vec3.Div(d_planar_sum, n)
		d_bulge := vec3.Mult(cell.normal, d_bulge_sum/n)

		nearby := 0
		for _, o := range c.cells {
			between := vec3.Subtract(cell.position, o.position)
			dist2 := between.LengthSqr()
			if dist2 < r2 {
				d_collision_sum = vec3.Add(d_collision_sum, vec3.Mult(between, (r2-dist2)/r2))
				nearby++
			}
		}
		d_collision := vec3.Div(d_collision_sum, float64(nearby)+n)

		p := vec3.Add(cell.position, vec3.Mult(d_spring, c.params.springFactor))
		p = vec3.Add(p, vec3.Mult(d_planar, c.params.planarFactor))
		p = vec3.Add(p, vec3.Mult(d_bulge, c.params.bulgeFactor))
		p = vec3.Add(p, vec3.Mult(d_collision, c.params.repulsionFactor))

	}
	c.updatePositionsAndFeed()
	t := time.Now()
	elapsed := t.Sub(start)
	var average time.Duration = time.Duration(elapsed.Nanoseconds() / int64(len(c.cells)))
	debug("Average time per cell " + average.String())
}

func (c *cellform) checkForSplits() {
	n := len(c.cells)
	for i := 0; i < n; i++ {
		if len(c.cells) >= c.maxCells {
			return
		}
		if c.cells[i].food > 1 {
			newCell := c.cells[i].Split()
			c.cells = append(c.cells, newCell)
		}
	}
}

func (c *cellform) updatePositionsAndFeed() {
	for i := 0; i < len(c.cells); i++ {
		cell := &c.cells[i]
		cell.position = cell.updatedPosition
		cell.food += c.params.feedRate
		debug(fmt.Sprintf("Food is %f", cell.food))
	}
}

func (c cellformParams) asString() string {
	return fmt.Sprintf("linkLength = %f\n"+
		"springFactor = %f\n"+
		"planarfactor = %f\n"+
		"bulgeFactor = %f\n"+
		"repulsionRange = %f\n"+
		"repulsionFactor = %f\n"+
		"feedRate = %f",
		c.linkLength, c.springFactor, c.planarFactor, c.bulgeFactor, c.repulsionRange, c.repulsionFactor, c.feedRate)
}
