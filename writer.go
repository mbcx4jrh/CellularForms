package main

import (
	"fmt"
)

func writePovRaySpheres(cells []cell) {
	for _, c := range cells {
		fmt.Println("sphere {")
		fmt.Println("    " + povrayVector(c.position) + ", 1")
		fmt.Println("    texture {GROWTH_T}")
		fmt.Println("}")
	}
}

func povrayVector(v Vector3) string {
	return fmt.Sprintf("<%v, %v, %v>", v.x, v.y, v.z)
}
