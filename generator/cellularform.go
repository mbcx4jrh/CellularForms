package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/downflux/go-geometry/nd/hypersphere"
	"github.com/downflux/go-kd/kd"
	"github.com/downflux/go-kd/point"
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
	cells    []Cell
	params   cellformParams
	rules    Rules
	maxCells int
}

func NewCellform(maxCells int, params cellformParams, rules Rules) *cellform {
	var cells []Cell = make([]Cell, maxCells)
	return &cellform{cells[0:0], params, rules, maxCells}
}

func (cf *cellform) Cell(i int) *Cell {
	return &cf.cells[i]
}

func (c *cellform) seedMesh(m mesh) {
	seedCells := importMesh(m)
	c.cells = append(c.cells, seedCells...)
}

func (cf *cellform) iterate() {
	cf.checkForSplits()
	debug(fmt.Sprintf("Iterating through %d cells", len(cf.cells)))
	link2 := cf.params.linkLength * cf.params.linkLength
	r2 := cf.params.repulsionRange * cf.params.repulsionRange
	start := time.Now()

	treePoints := make([]point.P, len(cf.cells))
	for i := range cf.cells {
		treePoints[i] = &cf.cells[i]
	}
	time_conv := time.Now()

	tree, err := kd.New(treePoints)
	if err != nil {
		log.Fatal("Error returned from building k-d tree", err)
	}
	time_tree := time.Now()
	for i := 0; i < len(cf.cells); i++ {
		cell := &(cf.cells[i])
		cell.age++
		cf.computeNormal(cell)
		d_spring_sum := vec3.Zero()
		d_planar_sum := vec3.Zero()
		d_bulge_sum := 0.0
		d_collision_sum := vec3.Zero()

		for _, n_idx := range cell.links {
			n := cf.Cell(n_idx)
			vectorBetween := vec3.Subtract(n.position, cell.position)
			vbNormalised := vec3.Normalize(vectorBetween)
			distance2 := vectorBetween.LengthSqr()
			d_spring_sum = vec3.Add(d_spring_sum,
				vec3.Subtract(vectorBetween,
					vec3.Mult(vbNormalised, cf.params.linkLength)))
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

		ns, err := kd.RadialFilter(
			tree,
			*hypersphere.New(cell.P(), cf.params.repulsionRange),
			func(p point.P) bool { return true })
		if err != nil {
			log.Fatal("Fucked up k-d tree search", err)
		}

		for _, o := range ns {

			between := vec3.Subtract(cell.position, o.(*Cell).position)
			dist2 := between.LengthSqr()
			if dist2 < r2 {
				d_collision_sum = vec3.Add(d_collision_sum, vec3.Mult(between, (r2-dist2)/r2))
			}
		}

		nearby := len(ns)

		d_collision := vec3.Div(d_collision_sum, float64(nearby)+n)

		p := vec3.Add(cell.position, vec3.Mult(d_spring, cf.params.springFactor))
		p = vec3.Add(p, vec3.Mult(d_planar, cf.params.planarFactor))
		p = vec3.Add(p, vec3.Mult(d_bulge, cf.params.bulgeFactor))
		p = vec3.Add(p, vec3.Mult(d_collision, cf.params.repulsionFactor))
		cell.updatedPosition = p
	}
	cf.updatePositionsAndFeed()
	t := time.Now()
	elapsed := t.Sub(start)
	var average time.Duration = time.Duration(elapsed.Nanoseconds() / int64(len(cf.cells)))
	debug("Average time per cell " + average.String())
	debug("Time per iteration    " + elapsed.String())
	debug("Interface conversion took " + time_conv.Sub(start).String())
	debug("Tree building took " + time_tree.Sub(time_conv).String())
}

func (c *cellform) checkForSplits() {
	n := len(c.cells)
	for i := 0; i < n; i++ {
		if len(c.cells) >= c.maxCells {
			return
		}
		if c.cells[i].food >= 1 {
			c.Split(i)
		}
	}
}

func (c *cellform) updatePositionsAndFeed() {
	for i := 0; i < len(c.cells); i++ {
		cell := &(c.cells[i])
		cell.position = cell.updatedPosition
		c.rules.feeder(cell, c.params.feedRate)
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
