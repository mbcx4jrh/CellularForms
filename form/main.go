package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type cell struct {
	x, y, z float64
}

var verbose bool = false
var headerFile string

func main() {

	var inputFilename string

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.StringVar(&headerFile, "h", "povray/default_render.pov", "The POV-Ray file to include at the beginning of the output file")
	flag.StringVar(&inputFilename, "i", "cell-00001.cf", "The file to process (output from the generator")
	flag.Parse()

	debug("Verbose is on")

	transform(inputFilename)
}

func transform(inputFilename string) {
	cells := readFromCF(inputFilename)

	writePovRayFile(cells, outputFilename(inputFilename))
}

func writePovRayFile(cells []cell, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	header, err := os.Open(headerFile)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, header)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range cells {
		file.WriteString("sphere {\n")
		file.WriteString(fmt.Sprintf("  <%v, %v, %v>, 1\n", c.x, c.y, c.z))
		file.WriteString("  texture { GROWTH_T }\n")
		file.WriteString(("}\n"))
	}
}

func outputFilename(inputFilename string) string {
	return inputFilename + ".pov"
}

func readFromCF(filename string) []cell {
	debugf("Input file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	cells := make([]cell, 500)

	for {
		line, err := reader.Read()
		switch {
		case err == nil:
			cells = append(cells, parseCell(line))
		case err == io.EOF:
			return cells
		case err != nil:
			log.Fatal(err)
		}
	}
}

func parseCell(csv []string) cell {
	return cell{getFloat(csv[0]),
		getFloat(csv[1]),
		getFloat(csv[2])}
}

func getFloat(p string) float64 {
	f, err := strconv.ParseFloat(p, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func debug(msg string) {
	if verbose {
		println(msg)
	}
}

func debugf(msg string, params ...interface{}) {
	debug(fmt.Sprintf(msg, params...))
}
