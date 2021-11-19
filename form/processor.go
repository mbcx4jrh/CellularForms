package main

import (
	"github.com/mbcx4jrh/vec3"
)

func centerOnBounds(cells []cell, s stats) {
	offset(cells, s.center)
}

func centerOnAverage(cells []cell, s stats) {
	offset(cells, s.average)
}

func offset(cells []cell, offset vec3.Vector3) {
	for i := range cells {
		cells[i].position = vec3.Subtract(cells[i].position, offset)
	}
}
