package flags

import (
	"flag"
)

type Flags struct {
	Update  bool
	Install bool
	Verbose bool
	Path    string
}

func NewArgs() *Flags {
	var update = false
	var install = false
	var verbose = false
	var path = ""

	flag.BoolVar(&verbose, "verbose", false, "")
	flag.StringVar(&path, "path", "", "")

	flag.Parse()
	var tail = flag.Args()

	if len(tail) >= 1 && tail[0] == "update" {
		update = true
	}

	if len(tail) >= 1 && tail[0] == "install" {
		install = true
	}

	return &Flags{
		Verbose: verbose,
		Update:  update,
		Install: install,
		Path:    path,
	}
}
