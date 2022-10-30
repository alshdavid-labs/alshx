package customflags

import (
	"flag"
	"os"
	"path/filepath"
)

type Flags struct {
	Find        map[string]bool
	Interactive bool
	Dir         string
	Version     bool
}

func NewFlags() *Flags {
	var findSlice ArrayFlags
	interactive := false
	version := false
	find := map[string]bool{}
	dir, _ := os.Getwd()

	flag.BoolVar(&interactive, "interactive", false, "")
	flag.BoolVar(&interactive, "i", false, "")

	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&version, "v", false, "")

	flag.Var(&findSlice, "target", "")
	flag.Var(&findSlice, "t", "")
	flag.Var(&findSlice, "find", "")
	flag.Var(&findSlice, "f", "")
	flag.Parse()
	tail := flag.Args()

	if len(tail) == 1 {
		dir, _ = filepath.Abs(flag.Args()[0])
	}

	for _, target := range findSlice {
		find[target] = true
	}

	return &Flags{
		Find:        find,
		Dir:         dir,
		Interactive: interactive,
		Version:     version,
	}
}
