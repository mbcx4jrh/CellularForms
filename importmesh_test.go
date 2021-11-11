package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertices(t *testing.T) {
	m := mesh{}
	m.vertices = append(m.vertices, Vector3{1.0, 1.0, 1.0}, Vector3{2.0, 2.0, 2.0}, Vector3{3.0, 3.0, 3.0})
	cells := importMesh(m)

	assert := assert.New(t)
	assert.Equal(3, len(cells))
	assert.Equal(Vector3{1.0, 1.0, 1.0}, cells[0].position)
	assert.Equal(Vector3{2.0, 2.0, 2.0}, cells[1].position)
	assert.Equal(Vector3{3.0, 3.0, 3.0}, cells[2].position)
	assert.Equal(0, len(cells[0].links))
}

//tests the 1-neighbourhood is linked ccw
func TestLinksNeighbourhood(t *testing.T) {
	m := mesh{}
	p1 := Vector3{0.0, 0.0, 0.0}
	p2 := Vector3{1.0, 0.0, 0.0}
	p3 := Vector3{0.0, 1.0, 0.0}
	p4 := Vector3{-1.0, -1.0, 0.0}
	m.vertices = append(m.vertices, p1, p2, p3, p4)
	m.triangles = append(m.triangles, [3]int{0, 1, 2}, [3]int{0, 2, 3}, [3]int{0, 3, 1})
	cells := importMesh(m)

	assert := assert.New(t)
	assert.Equal(4, len(cells))
	assert.Equal(3, len(cells[0].links))
	assert.Equal(p2, cells[0].links[0].position)
	assert.Equal(p3, cells[0].links[1].position)
	assert.Equal(p4, cells[0].links[2].position)
}
