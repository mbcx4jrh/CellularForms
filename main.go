package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var verbose bool

func main() {

	var iterations int
	var debugFreq int
	var maxCells int
	var outputFreq int
	var folder string
	var filePrefix string
	var form *cellform

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.IntVar(&iterations, "i", 100, "Number of iterations to compute")
	flag.IntVar(&debugFreq, "df", 1, "If verbose set, iteration debug will be output after this number of iterations")
	flag.IntVar(&maxCells, "n", 1000, "Maximum number of cells that can be generated")
	flag.IntVar(&outputFreq, "of", 10, "the frequency at whoh an output file is generated")
	flag.StringVar(&folder, "o", "output", "The folder name for the output files")
	flag.StringVar(&filePrefix, "f", "cell-", "The file name prefix to use for output files")
	flag.Parse()

	debug("Verbose output is on")
	debug("Initialising cell array to maximum of " + strconv.Itoa(maxCells) + "...")

	form = NewCellform(maxCells)
	form.seedMesh(isocahedron())
	debug("Initial seed mesh contains " + strconv.Itoa(len(form.cells)) + " cells")

	writer := NewCellWriter(filePrefix, folder)
	err := writer.initialise()
	if err != nil {
		log.Fatal(err)
	}

	debug("Running for " + strconv.Itoa(iterations) + " iterations")
	for i := 0; i < iterations; i++ {
		if i%debugFreq == 0 {
			debug(fmt.Sprintf("Iteration %d", i))
		}
		if i%outputFreq == 0 {
			writer.writeNextFile(form.cells)
		}
		form.iterate()
	}

}

func debug(message string) {
	if verbose {
		println(message)
	}
}
