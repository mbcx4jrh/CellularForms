package main

import "github.com/mbcx4jrh/vec3"

type cell struct {
	index           int
	position        vec3.Vector3
	updatedPosition vec3.Vector3
	normal          vec3.Vector3
	links           []cell
}

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
