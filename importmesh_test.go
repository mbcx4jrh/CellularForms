package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/mbcx4jrh/vec3"
)

func TestVertices(t *testing.T) {
	m := mesh{}
	m.vertices = append(m.vertices, Vector3{X: 1.0, Y: 1.0, Z: 1.0}, Vector3{X: 2.0, Y: 2.0, Z: 2.0}, Vector3{X: 3.0, Y: 3.0, Z: 3.0})
	cells := importMesh(m)

	assert := assert.New(t)
	assert.Equal(3, len(cells))
	assert.Equal(Vector3{X: 1.0, Y: 1.0, Z: 1.0}, cells[0].position)
	assert.Equal(Vector3{X: 2.0, Y: 2.0, Z: 2.0}, cells[1].position)
	assert.Equal(Vector3{X: 3.0, Y: 3.0, Z: 3.0}, cells[2].position)
	assert.Equal(0, len(cells[0].links))
}

//tests the 1-neighbourhood is linked ccw
func TestLinksNeighbourhood(t *testing.T) {
	m := mesh{}
	p1 := Vector3{X: 0.0, Y: 0.0, Z: 0.0}
	p2 := Vector3{X: 1.0, Y: 0.0, Z: 0.0}
	p3 := Vector3{X: 0.0, Y: 1.0, Z: 0.0}
	p4 := Vector3{X: -1.0, Y: -1.0, Z: 0.0}
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

func TestIsocahedron(t *testing.T) {
	m := isocahedron()
	cells := importMesh(m)
	for i := 0; i < len(cells); i++ {
		assert.Greater(t, len(cells[i].links), 0)
	}
	for _, c := range cells {
		assert.Greater(t, len(c.links), 0)
	}
	debug(fmt.Sprintf("cell 4 is id %d", cells[0].links[4].id))
	assert.Greater(t, len(cells[0].links[4].links), 0)
}
