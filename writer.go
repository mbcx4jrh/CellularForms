package main

import "fmt"

func writePovRaySpheres(cells []cell) {
	for i := 0; i < len(cells); i++ {
		fmt.Println("sphere {")
		fmt.Println("    <0, 0, 0>, 1")
		fmt.Println("    texture {GROWTH_T}")
		fmt.Println("}")
	}
}
