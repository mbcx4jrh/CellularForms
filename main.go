package main

import (
	"flag"
	"strconv"
)

var verbose bool

func main() {

	var iterations int
	var maxCells int
	var form *cellform

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.IntVar(&iterations, "i", 100, "Number of iterations to compute")
	flag.IntVar(&maxCells, "n", 1000, "Maximum number of cells that can be generated")

	flag.Parse()

	debug("Verbose output is on")
	debug("Initialising cell array to maximum of " + strconv.Itoa(maxCells) + "...")

	form = createCellsStructure(maxCells)
	form.seedMesh(isocahedron())
	debug("Initial seed mesh contains " + strconv.Itoa(len(form.cells)) + " cells")

	debug("Running for " + strconv.Itoa(iterations) + " iterations")

	debug("<--POVRAY START-->")
	writePovRaySpheres(form.cells)
	debug("<--POVRAY END-->")
}

func debug(message string) {
	if verbose {
		println(message)
	}
}
