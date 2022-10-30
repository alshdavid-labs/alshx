package customflags

import (
	"flag"
)

type Flags struct {
	Action  string
	Length  int64
	Version bool
}

func NewFlags() *Flags {
	var action string = ""
	var length int64 = 0
	var version bool = false

	flag.Int64Var(&length, "length", 10, "")
	flag.Int64Var(&length, "l", 10, "")

	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&version, "v", false, "")

	flag.Parse()
	tail := flag.Args()

	if len(tail) == 1 {
		action = tail[0]
	}

	return &Flags{
		Action:  action,
		Length:  length,
		Version: version,
	}
}
