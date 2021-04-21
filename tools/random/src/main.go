package main

import (
	"fmt"
	"os"
	"rrm/src/customflags"
)

func main() {
	args := customflags.NewFlags()

	if args.Version {
		fmt.Println("Generate random things")
		fmt.Println("EG:")
		fmt.Println("  random string")
		fmt.Println("  random --length 20 string")
		fmt.Println("")
		fmt.Println("  random number")
		fmt.Println("  random --length 20 number")
		fmt.Println("  random guid")
		os.Exit(0)
	}
}
