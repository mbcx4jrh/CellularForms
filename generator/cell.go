package main

import (
	"fmt"
	"math"

	"github.com/downflux/go-geometry/nd/vector"
	"github.com/mbcx4jrh/vec3"
)

type Cell struct {
	id              uint64
	position        vec3.Vector3
	updatedPosition vec3.Vector3
	normal          vec3.Vector3
	food            float64
	trait           int64
	age             float64
	links           []int
}

var nextCellId uint64

func (c *Cell) P() vector.V {
	return *vector.New(c.position.X, c.position.Y, c.position.Z)
}

func (cf *cellform) computeNormal(c *Cell) {
	sum := vec3.Zero()
	for i, n := range c.links {
		a := vec3.Subtract(cf.Cell(n).position, c.position)
		b := vec3.Subtract(cf.Cell(c.links[(i+1)%len(c.links)]).position, c.position)
		sum = vec3.Add(sum, vec3.Cross(a, b))
	}
	newNormal := vec3.Normalize(sum)

	if vec3.Dot(c.normal, newNormal) < 0 {
		newNormal = vec3.Subtract(vec3.Zero(), newNormal)
	}

	c.normal = newNormal
}

func (cf *cellform) Split(c_idx int) {
	c := cf.Cell(c_idx)

	cf.cells = append(cf.cells, NewCell(c.position, c.normal))
	d_idx := len(cf.cells) - 1
	daughter := &cf.cells[d_idx]

	daughter.position = c.position
	daughter.normal = c.normal
	daughter.trait = c.trait
	n := len(c.links)
	if n == 0 {
		return
	}

	//find nearest neighbour
	nearest_d := math.MaxFloat64
	nearest := -1
	for i, neighbour := range c.links {
		d := vec3.Subtract(cf.Cell(neighbour).position, c.position).LengthSqr()
		if d < nearest_d {
			nearest_d = d
			nearest = i
		}
	}
	opposite := (nearest + n/2) % n
	//parent links
	newLinks := []int{}
	for i := nearest; i != (opposite+1)%n; i = (i + 1) % n {
		newLinks = append(newLinks, c.links[i])
	}
	newLinks = append(newLinks, d_idx)

	//daughter links
	daughter.links = append(daughter.links, c.links[opposite])
	for i := (opposite + 1) % n; i != nearest; i = (i + 1) % n {
		daughter.links = append(daughter.links, c.links[i])

		cf.Cell(c.links[i]).replaceLink(c_idx, d_idx)
	}
	cf.Cell(c.links[nearest]).addAfter(c_idx, d_idx)
	cf.Cell(c.links[opposite]).addBefore(c_idx, d_idx)
	daughter.links = append(daughter.links, c.links[nearest])
	daughter.links = append(daughter.links, c_idx)

	c.links = newLinks

	cf.computeNewPosition(c)
	cf.computeNewPosition(daughter)
	c.food -= 1

	//CellReport("after split - parent ", c)
	//CellReport("after split - child  ", daughter)
}

func NewCell(position, normal vec3.Vector3) Cell {
	c := Cell{nextCellId, position, position, normal, 0, 0, 0, []int{}}
	//c := Cell{}
	//c.id = nextCellId
	//c.position = position
	//c.normal = position
	nextCellId++
	return c
}

func (cf *cellform) computeNewPosition(c *Cell) {
	p := vec3.Zero()
	for _, n := range c.links {
		p = vec3.Add(p, cf.Cell(n).position)
	}
	count := float64(len(c.links))
	c.position = vec3.Div(p, count)
}

func (c *Cell) addAfter(after, newCell int) {
	i := indexOf(c.links, after)
	c.links = insert(c.links, newCell, i+1)
}

func (c *Cell) addBefore(before, newCell int) {
	i := indexOf(c.links, before)
	c.links = insert(c.links, newCell, i)
}

func (c *Cell) replaceLink(old, newCell int) {
	i := indexOf(c.links, old)
	c.links[i] = newCell
}

func indexOf(a []int, v int) int {
	for i, c := range a {
		if c == v {
			return i
		}
	}
	debug(fmt.Sprintf("Couldnt find cell %d in %d links", v, len(a)))
	return -1
}

func insert(a []int, c int, i int) []int {
	if len(a) == i {
		return append(a, c)
	}
	return append(a[:i], append([]int{c}, a[i:]...)...)
}
