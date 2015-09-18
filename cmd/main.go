package main

import (
	"github.com/washtubs/stratifier"
)

func main() {
	// TODO: make this just take a file as the first arg and print the default stratification
	stratifier.ReadJson("test.json")

	// LONGTERM TODO:
	// do I need a nice arg parsing lib?
	//    should probably use "flag" package
	// TODO: parse args
	// NOTE: should be super configurable
}
