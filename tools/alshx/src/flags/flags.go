package flags

import (
	"flag"
)

type Flags struct {
	Update  bool
	Install bool
	Verbose bool
	Version bool
	Dry     bool
	Path    string
}

func NewArgs() *Flags {
	var update = false
	var install = false
	var verbose = false
	var version = false
	var dry = false
	var path = ""

	flag.BoolVar(&verbose, "verbose", false, "")
	flag.BoolVar(&version, "version", false, "")
	flag.StringVar(&path, "path", "", "")
	flag.BoolVar(&dry, "dry", false, "")

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
		Version: version,
		Path:    path,
		Dry:     dry,
	}
}
