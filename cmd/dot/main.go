package main

import (
	"flag"
	"fmt"

	"github.com/awakot/danceOnTerm/internal/app"
)

const (
	version = "1.0.0"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}
	app.StartApp()
}
