package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mbcx4jrh/vec3"
)

type cell struct {
	position vec3.Vector3
	age      float64
	trait    int64
}

var verbose bool
var headerFile string
var offsetAverage bool
var scale float64
var traitsAsTextures bool
var slice float64

func main() {

	var inputFilename string

	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.BoolVar(&offsetAverage, "ca", false, "Center form on average position rather than midpoint between bounds")
	flag.StringVar(&headerFile, "h", "povray/default_render.pov", "The POV-Ray file to include at the beginning of the output file")
	flag.StringVar(&inputFilename, "i", "cell-00001.cf", "The file to process (output from the generator")
	flag.Float64Var(&scale, "s", 1, "Scaling factor for resultant image")
	flag.BoolVar(&traitsAsTextures, "traits", false, "traits are inherited from parents and decide the texture of the cell")
	flag.Float64Var(&slice, "slice", -100000000, "Removes all cells below this z coordinate")
	flag.Parse()

	debug("Verbose is on")

	transform(inputFilename)
}

func transform(inputFilename string) {
	cells := readFromCF(inputFilename)

	stats := GetStats(cells)

	if offsetAverage {
		centerOnAverage(cells, stats)
	} else {
		centerOnBounds(cells, stats)
	}

	writePovRayFile(cells, outputFilename(inputFilename), stats)
}

func writePovRayFile(cells []cell, filename string, s stats) {
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
	writeCameraAndlight(file, s)

	for _, c := range cells {
		if c.position.Z < slice {
			continue
		}
		file.WriteString("sphere {\n")
		file.WriteString(fmt.Sprintf("  <%v, %v, %v>, 1\n", c.position.X, c.position.Y, c.position.Z))
		if traitsAsTextures {
			file.WriteString(fmt.Sprintf("  texture { GROWTH_T_%v }\n", c.trait))
		} else {
			file.WriteString("  texture { GROWTH_T }\n")
		}
		file.WriteString(("}\n"))
	}
}

func writeCameraAndlight(file *os.File, s stats) {
	camera := vec3.New(0, 2, -5)
	light1 := vec3.New(2, 4, -3)
	light2 := vec3.New(-2, 4, -5)

	width := (s.max.X - s.min.X + 1.0) / scale // the 1,0 is the width of a sphere
	debugf("Using scaled width of %v", width)
	camera = vec3.Mult(camera, width)
	light1 = vec3.Mult(light1, width)
	light2 = vec3.Mult(light2, width)

	file.WriteString("camera {\n")
	file.WriteString(fmt.Sprintf("  location <%v, %v, %v>\n", camera.X, camera.Y, camera.Z))
	file.WriteString("  look_at <0, 0, 1>\n")
	file.WriteString("}\n")

	writeLight(file, light1)
	writeLight(file, light2)
}

func writeLight(file *os.File, p vec3.Vector3) {
	file.WriteString("light_source {\n")
	file.WriteString(fmt.Sprintf("  <%v, %v, %v> colour White\n", p.X, p.Y, p.Z))
	file.WriteString("}\n")
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

	cells := make([]cell, 0, 500)

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
	position := vec3.New(getFloat(csv[0]), getFloat(csv[1]), getFloat(csv[2]))
	age := getFloat(csv[3])
	trait := getInt(csv[4])
	return cell{position, age, trait}
}

func getInt(p string) int64 {
	i, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
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
