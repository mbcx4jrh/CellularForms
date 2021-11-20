package main

import (
	"github.com/downflux/go-geometry/nd/vector"
)

type TreePoint struct {
	p         vector.V
	cellIndex int
}

func (tp TreePoint) P() vector.V {
	return tp.p
}
