package handler_number

import "flag"

func parseArgs(a []string) *args {
	f := flag.NewFlagSet("", flag.ExitOnError)

	length := 0

	f.IntVar(&length, "length", 0, "Length of string")
	f.IntVar(&length, "l", 0, "Length of string")

	f.Parse(a)

	if length == 0 {
		length = 20
	}

	return &args{
		length: length,
	}
}
