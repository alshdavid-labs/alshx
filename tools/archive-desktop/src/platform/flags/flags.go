package flags

import (
	"flag"
)

type Flags struct {
	Dry     bool
	Version bool
	Verbose bool
}

func NewArgs() *Flags {
	dry := false
	version := false
	verbose := false

	flag.BoolVar(&dry, "dry", false, "")
	flag.BoolVar(&dry, "d", false, "")

	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&version, "v", false, "")

	flag.BoolVar(&verbose, "verbose", false, "")

	flag.Parse()
	// tail := flag.Args()

	return &Flags{
		Dry:     dry,
		Version: version,
		Verbose: verbose,
	}
}
