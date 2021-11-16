package main

import (
	"fmt"
	"log"
)

//some debug functions

func CellReport(msg string, c *cell) {
	format := msg + " Cell %d, links to %v"
	debug(fmt.Sprintf(format, c.id, linkIds((c.links))))
}

func linkIds(links []*cell) []int {

	ids := []int{}
	for _, c := range links {
		ids = append(ids, int(c.id))
	}
	return ids
}

//check all links are returned - used for debug
func ValidateLinks(c *cell, msg string) {
	for _, n := range c.links {
		q := indexOf(n.links, c)
		if q == -1 {
			CellReport("in validate", c)
			CellReport("link is ", n)
			log.Fatal(fmt.Sprintf(msg+" Fucked up links - can't find return from %d to %d", n.id, c.id))
		}
	}
}

func ValidateLinksGlobally(cells []cell, msg string) {
	for _, c := range cells {
		ValidateLinks(&c, msg)
	}
}
