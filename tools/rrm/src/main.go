package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"rrm/src/customflags"
)

func main() {
	args := customflags.NewFlags()

	if args.Version {
		fmt.Println("Recursive Remove, probably the latest verison")
		os.Exit(0)
	}
	crawl(args.Dir, args.Find, args.Dry)
}

func crawl(dir string, targets map[string]bool, dry bool) {
	fileInfo, _ := os.Stat(dir)
	if !fileInfo.IsDir() {
		fmt.Println("Target is not a folder")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if targets[name] {
			if !dry {
				fmt.Println("Delete:", path)
				os.RemoveAll(path)
			} else {
				fmt.Println("DryDelete:", path)
			}
			continue
		}

		if file.IsDir() {
			crawl(path, targets, dry)
			continue
		}
	}
}
