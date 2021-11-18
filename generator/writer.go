package main

import (
	"fmt"
	"log"
	"os"
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

	f, err := os.Create(fmt.Sprintf(c.folder+"/"+c.filePrefx+"%05d.cf", c.nextFile))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, c := range cells {
		writeCell(f, c)
	}
	c.nextFile++
}

func writeCell(f *os.File, c Cell) {
	f.WriteString(fmt.Sprintf("%v,%v,%v\n", c.position.X, c.position.Y, c.position.Z))
}
