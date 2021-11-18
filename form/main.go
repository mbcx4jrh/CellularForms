package main

import "flag"

var verbose bool = false

func main() {
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()
}
