package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alshdavid/alshx/tools/rrm/platform/customflags"
)

func main() {
	args := customflags.NewFlags()

	if args.Version {
		fmt.Println("Recursive Remove, probably the latest verison")
		os.Exit(0)
	}

	if len(args.Find) == 0 {
		fmt.Println("No Arguments Specified")
		fmt.Println("")
		fmt.Println("Example:")
		fmt.Println("  rrm -f node_modules -f .DS_Store ./")
		fmt.Println("")
		fmt.Println("This will find any file or folder with")
		fmt.Println("the name s\"node_modeles\" and \".DS_Store\"")
		fmt.Println("in the current directory")

		os.Exit(0)
	}

	fmt.Printf("Looking Within:\n")
	fmt.Printf("  %s\n\n", args.Dir)

	for dir := range args.Find {
		fmt.Printf("Seeking:\n")
		fmt.Printf("  - %s\n", dir)
	}

	foundItems := map[string]bool{}

	fmt.Printf("\nFound:\n")
	scan(foundItems, args.Find, args.Dir)

	if len(foundItems) == 0 {
		fmt.Printf("  Nothing to delete\n")
		os.Exit(0)
	}

	if args.Interactive {
		fmt.Printf("\nOk to delete? [Y/n] ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		answer := strings.Trim(strings.Replace(strings.ToLower(input), "\n", "", -1), "")

		if answer == "n" || answer == "no" {
			os.Exit(0)
			return
		}
	}

	fmt.Printf("\nRemoved:\n")

	for pathToDelete := range foundItems {
		fmt.Printf("  %s\n", pathToDelete)
		os.RemoveAll(pathToDelete)
	}
}

func scan(foundItems map[string]bool, lookupNames map[string]bool, lookupDirectory string) {
	files, err := ioutil.ReadDir(lookupDirectory)
	if err != nil {
		fmt.Printf("  SKIP: %s\n", lookupDirectory)
		return
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(lookupDirectory, name)

		if lookupNames[name] {
			foundItems[path] = true
			fmt.Printf("  HIT:  %s\n", path)
			continue
		}

		if file.IsDir() {
			scan(foundItems, lookupNames, path)
		}
	}
}
