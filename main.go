package main

import (
	"flag"
	"fmt"
	"strconv"
)

var verbose bool

func main() {

	var iterations int
	var iterationFreq int
	var maxCells int
	var form *cellform

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.IntVar(&iterations, "i", 100, "Number of iterations to compute")
	flag.IntVar(&iterationFreq, "if", 1, "If verbose set, iteration debug will be output after this number of iterations")
	flag.IntVar(&maxCells, "n", 1000, "Maximum number of cells that can be generated")

	flag.Parse()

	debug("Verbose output is on")
	debug("Initialising cell array to maximum of " + strconv.Itoa(maxCells) + "...")

	form = NewCellform(maxCells)
	form.seedMesh(isocahedron())
	debug("Initial seed mesh contains " + strconv.Itoa(len(form.cells)) + " cells")

	debug("Running for " + strconv.Itoa(iterations) + " iterations")
	for i := 0; i < iterations; i++ {
		if i%iterationFreq == 0 {
			debug(fmt.Sprintf("Iteration %d", i))
		}
		form.iterate()
	}

	debug("<--POVRAY START-->")
	writePovRaySpheres(form.cells)
	debug("<--POVRAY END-->")
}

func debug(message string) {
	if verbose {
		println("// " + message)
	}
}
