package main

import (
	"flag"

	"github.com/dkvilo/andromeda/sandbox"
)

var flagvar string

func init() {
	flag.StringVar(&flagvar, "scene", "3d", "specify renderer type")
}

func main() {
	flag.Parse()
	if flagvar == "3d" {
		sandbox.Scene3d.Andromeda.Run()
	} else if flagvar == "blank" {
		sandbox.Blank.Andromeda.Run()
	}
}
