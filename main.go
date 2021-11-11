package main

import (
	"flag"
)

var verbose bool

func main() {

	var noOfIterations int
	var maxCells int
	var cells []cell

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.IntVar(&noOfIterations, "i", 100, "Number of iterations to compute")
	flag.IntVar(&maxCells, "n", 1000, "Maximum number of cells that can be generated")

	flag.Parse()

	cells = createCellsStructure(maxCells)

	debug("<--POVRAY START-->")
	writePovRaySpheres(cells)
	debug("<--POVRAY END-->")
}

func debug(message string) {
	if verbose {
		println(message)
	}
}
