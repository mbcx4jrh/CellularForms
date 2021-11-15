package main

import (
	"fmt"
	"math"

	"github.com/mbcx4jrh/vec3"
)

type cell struct {
	id              uint64
	position        vec3.Vector3
	updatedPosition vec3.Vector3
	normal          vec3.Vector3
	food            float64
	links           []*cell
}

var nextCellId uint64

func (c *cell) computeNormal() {
	sum := vec3.Zero()
	for i, n := range c.links {
		a := vec3.Subtract(n.position, c.position)
		b := vec3.Subtract(c.links[(i+1)%len(c.links)].position, c.position)
		//b := NewSubtract(&c.links[(i+1)%len(c.links)].position, &c.position)
		sum = vec3.Add(sum, vec3.Cross(a, b))
	}
	newNormal := vec3.Normalize(sum)

	if vec3.Dot(c.normal, newNormal) < 0 {
		newNormal = vec3.Subtract(vec3.Zero(), newNormal)
	}

	c.normal = newNormal
}

func (c *cell) Split(daughter *cell) {
	//debug(fmt.Sprintf("Splitting cell %d with %d links", c.id, len(c.links)))
	//daughter := NewCell(c.position, c.normal)
	daughter.position = c.position
	daughter.normal = c.normal
	n := len(c.links)
	if n == 0 {
		return
	}
	//find nearest neighbour
	nearest_d := math.MaxFloat64
	nearest := -1
	for i, neighbour := range c.links {
		d := vec3.Subtract(neighbour.position, c.position).LengthSqr()
		if d < nearest_d {
			nearest_d = d
			nearest = i
		}
	}
	opposite := (nearest + n/2) % n

	//parent links
	newLinks := []*cell{}
	for i := nearest; i != (opposite+1)%n; i = (i + 1) % n {
		newLinks = append(newLinks, c.links[i])
	}
	newLinks = append(newLinks, daughter)

	//daughter links
	daughter.links = append(daughter.links, c.links[opposite])
	for i := (opposite + 1) % n; i != nearest; i = (i + 1) % n {
		daughter.links = append(daughter.links, c.links[i])
		//debug(fmt.Sprintf("About to replace link for %d in cell %d (%d links)", c.id, c.links[i].id, len(c.links[i].links)))
		c.links[i].replaceLink(c, daughter)
	}
	c.links[nearest].addAfter(c, daughter)
	c.links[opposite].addBefore(c, daughter)
	daughter.links = append(daughter.links, c.links[nearest])
	daughter.links = append(daughter.links, c)

	c.links = newLinks

	c.computeNewPosition()
	daughter.computeNewPosition()
	c.food -= 1
}

func NewCell(position, normal vec3.Vector3) cell {
	c := cell{nextCellId, position, position, normal, 0, []*cell{}}
	nextCellId++
	return c
}

func (c *cell) computeNewPosition() {
	p := c.position
	for _, n := range c.links {
		p = vec3.Add(p, n.position)
	}
	count := float64(len(c.links)) + 1
	c.position = vec3.Div(p, count)
}

func (c *cell) addAfter(after, newCell *cell) {
	i := indexOf(c.links, after)
	c.links = insert(c.links, newCell, i+1)
}

func (c *cell) addBefore(before, newCell *cell) {
	i := indexOf(c.links, before)
	c.links = insert(c.links, newCell, i)
}

func (c *cell) replaceLink(old, newCell *cell) {
	i := indexOf(c.links, old)
	c.links[i] = newCell
}

func indexOf(a []*cell, v *cell) int {
	for i, c := range a {
		if c.id == v.id {
			return i
		}
	}
	debug(fmt.Sprintf("Couldnt find cell %d in %d links", v.id, len(a)))
	return -1
}

func insert(a []*cell, c *cell, i int) []*cell {
	if len(a) == i {
		return append(a, c)
	}
	return append(a[:i], append([]*cell{c}, a[i:]...)...)
}
