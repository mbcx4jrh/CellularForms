package main

import (
	"testing"

	"github.com/mbcx4jrh/vec3"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	assert := assert.New(t)

	nextCellId = 0
	c1 := MakeZeroCell()
	c2 := MakeZeroCell()
	c3 := MakeZeroCell()
	c4 := MakeZeroCell()

	r := MakeZeroCell()

	c1.links = append(c1.links, c2, c3, c4)
	assert.Equal(3, len(c1.links))
	assert.Equal(c3.id, c1.links[1].id)
	assert.NotEqual(c3.id, r.id)

	c1.replaceLink(c3, r)

	assert.Equal(3, len(c1.links))
	assert.Equal(r.id, c1.links[1].id)
}

func MakeZeroCell() *cell {
	c := NewCell(vec3.Zero(), vec3.Zero())
	return &c
}

func TestPointersForLinks(t *testing.T) {

	p1 := vec3.Zero()

	c1 := NewCell(p1, p1)
	c2 := NewCell(p1, p1)

	c1.links = append(c1.links, &c2)

	p2 := vec3.New(1, 1, 1)
	c2.position = p2

	assert.Equal(t, p2, c1.links[0].position)
}

func TestPointerForSliceOfCells(t *testing.T) {
	cells := []cell{}
	c := NewCell(vec3.Zero(), vec3.Zero())
	cells = append(cells, c)

	cells[0].food = 1

	assert.Equal(t, 1.0, cells[0].food)

	aCell := &cells[0]

	aCell.food += 0.1

	assert.Equal(t, 1.1, cells[0].food)

}

func TestSimple4Split(t *testing.T) {
	nextCellId = 0
	p1 := vec3.Zero()
	p2 := vec3.New(-1, 1, 0)
	p3 := vec3.New(-1, -1, 0)
	p4 := vec3.New(1, -1, 0)
	p5 := vec3.New(1, 1, 0)

	c1 := NewCell(p1, p1)
	c2 := NewCell(p2, p2)
	c3 := NewCell(p3, p3)
	c4 := NewCell(p4, p4)
	c5 := NewCell(p5, p5)

	c1.links = append(c1.links, &c2, &c3, &c4, &c5)
	c2.links = append(c2.links, &c3, &c1, &c5)
	c3.links = append(c3.links, &c4, &c1, &c2)
	c4.links = append(c4.links, &c5, &c1, &c3)
	c5.links = append(c5.links, &c2, &c1, &c4)

	d := NewCell(p1, p1)

	c1.Split(&d)

	assert := assert.New(t)
	assert.Equal(4, len(c1.links))
	assert.Equal(4, len(d.links))

	assert.Equal(4, len(c2.links))
	assert.Equal(3, len(c3.links))
	assert.Equal(4, len(c4.links))
	assert.Equal(3, len(c5.links))
}

func TestSimple5Split(t *testing.T) {
	nextCellId = 0
	p0 := vec3.Zero()
	p1 := vec3.New(0, 1, 0)
	p2 := vec3.New(-1, 1, 0)
	p3 := vec3.New(-1, -1, 0)
	p4 := vec3.New(1, -1, 0)
	p5 := vec3.New(1, 1, 0)

	c0 := NewCell(p0, p0)
	c1 := NewCell(p1, p1)
	c2 := NewCell(p2, p2)
	c3 := NewCell(p3, p3)
	c4 := NewCell(p4, p4)
	c5 := NewCell(p5, p5)

	c0.links = append(c0.links, &c1, &c2, &c3, &c4, &c5)
	c1.links = append(c1.links, &c2, &c0, &c5)
	c2.links = append(c2.links, &c3, &c0, &c1)
	c3.links = append(c3.links, &c4, &c0, &c2)
	c4.links = append(c4.links, &c5, &c0, &c3)
	c5.links = append(c5.links, &c1, &c0, &c4)

	d := NewCell(p1, p1)
	ad := &d
	c1.Split(ad)

	assert := assert.New(t)
	assert.Equal(3, len(c1.links))
	assert.Equal(4, len(d.links))

	assert.Equal(4, len(c2.links))
	assert.Equal(3, len(c3.links))
	assert.Equal(3, len(c4.links))
	assert.Equal(3, len(c5.links))

	d2 := NewCell(vec3.Zero(), vec3.Zero())
	d.Split(&d2)
	p := vec3.Vector3{X: 100, Y: 100, Z: 100}
	c0.position = p
	c1.position = p
	c2.position = p
	c3.position = p
	c4.position = p
	c5.position = p
	d.position = p
	d2.position = p
	//ad.position = vec3.Zero()
	debug("Will this fail")
}
