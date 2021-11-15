package main

import (
	"testing"

	"github.com/mbcx4jrh/vec3"
	"github.com/stretchr/testify/assert"
)

func TestPointersForLinks(t *testing.T) {

	p1 := vec3.Zero()

	c1 := NewCell(p1, p1)
	c2 := NewCell(p1, p1)

	c1.links = append(c1.links, &c2)

	p2 := vec3.New(1, 1, 1)
	c2.position = p2

	assert.Equal(t, p2, c1.links[0].position)
}
