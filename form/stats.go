package main

import (
	"github.com/mbcx4jrh/vec3"
)

type stats struct {
	min, max vec3.Vector3
	center   vec3.Vector3
	average  vec3.Vector3
}

func GetStats(cells []cell) stats {
	if len(cells) == 0 {
		return stats{}
	}
	min := cells[0].position
	max := cells[0].position
	average := vec3.Zero()
	for _, c := range cells {
		checkBounds(&min, &max, c)
		average = vec3.Add(average, c.position)
	}

	center := vec3.Mult(vec3.Add(min, max), 0.5)
	average = vec3.Div(average, float64(len(cells)))

	s := stats{min, max, center, average}

	return s
}

func checkBounds(min, max *vec3.Vector3, c cell) {
	p := c.position
	if p.X < min.X {
		min.X = p.X
	}
	if p.Y < min.Y {
		min.Y = p.Y
	}
	if p.Z < min.Z {
		min.Z = p.Z
	}
	if p.X > max.X {
		max.X = p.X
	}
	if p.Y > max.Y {
		max.Y = p.Y
	}
	if p.Z > max.Z {
		max.Z = p.Z
	}
}
