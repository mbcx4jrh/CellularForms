package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mbcx4jrh/vec3"
)

type cellWriter struct {
	filePrefx string
	folder    string
	nextFile  int
}

func NewCellWriter(filePrefix string, folder string) cellWriter {
	return cellWriter{filePrefix, folder, 1}
}

func (c cellWriter) initialise() error {
	_, err := os.Stat(c.folder)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(c.folder, 0755)
		if errDir != nil {
			return errDir
		}
		return nil
	}
	return err
}

func (c *cellWriter) writeNextFile(cells []Cell) {

	f, err := os.Create(fmt.Sprintf(c.folder+"/"+c.filePrefx+"%05d.pov", c.nextFile))
	if err != nil {
		log.Fatal(err)
	}
	copyDefault(f, "povray/default_render.pov")
	defer f.Close()
	for _, c := range cells {
		writeCell(f, c)
	}
	c.nextFile++
}

func writeCell(f *os.File, c Cell) {
	f.WriteString("sphere {\n")
	f.WriteString("    " + povrayVector(c.position) + ", 1\n")
	f.WriteString("    texture {GROWTH_T}\n")
	f.WriteString("}\n")
}

func copyDefault(f *os.File, defaultFile string) {
	content, err := os.ReadFile(defaultFile)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(content)
}

func povrayVector(v vec3.Vector3) string {
	return fmt.Sprintf("<%v, %v, %v>", v.X, v.Y, v.Z)
}
