package main

import (
	"fmt"
	"log"
)

//some debug functions

func CellReport(msg string, c *cell) {
	format := msg + " Cell %d, links to %v"
	debug(fmt.Sprintf(format, c.id, c.links))
}

//check all links are returned - used for debug
func (cf *cellform) ValidateLinks(c_idx int, msg string) {
	c := cf.Cell(c_idx)
	for _, n_idx := range c.links {
		n := cf.Cell(n_idx)
		q := indexOf(n.links, c_idx)
		if q == -1 {
			CellReport("in validate", c)
			CellReport("link is ", n)
			log.Fatal(fmt.Sprintf(msg+" Fucked up links - can't find return from %d to %d", n.id, c.id))
		}
	}
}

func (cf *cellform) ValidateLinksGlobally(msg string) {
	for i := range cf.cells {
		cf.ValidateLinks(i, msg)
	}
}
