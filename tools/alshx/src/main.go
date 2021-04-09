package main

import (
	"archivedesktop/src/archive"
	"archivedesktop/src/download"
	"archivedesktop/src/flags"
	"archivedesktop/src/github"
	"archivedesktop/src/logging"
	"archivedesktop/src/paths"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var err error
	var args = flags.NewArgs()
	var logger = logging.NewLogger(args.Verbose)

	if !args.Update && !args.Install {
		fmt.Println("Alshx Command Line Utilities")
		os.Exit(0)
	}

	var alshxPath = args.Path
	if alshxPath == "" {
		alshxPath = filepath.Join(paths.GetHomePath(), ".local", "alshx")
	}

	fmt.Println("Alshx Command Line Utilities")
	fmt.Println("Folder to install to:", alshxPath)

	tempDir, err := ioutil.TempDir("", "alshx")
	if err != nil {
		logger.Error("ERROR: Unable to create temp directory")
		os.Exit(1)
	}

	var archivePath = filepath.Join(tempDir, github.GetArchiveName())

	defer os.RemoveAll(tempDir)

	logger.Log("---")

	logger.Log("25% Downloading new release")
	err = download.File(github.GetReleaseUrl(), archivePath)
	if err != nil {
		logger.Error("ERROR: Unable to download archive")
		os.Exit(1)
	}

	logger.Log("50% Removing existing binaries")
	binaries, _ := ioutil.ReadDir(alshxPath)
	for _, file := range binaries {
		os.RemoveAll(filepath.Join(alshxPath, file.Name()))
	}

	logger.Log("75% Unpacking binaries")
	os.MkdirAll(alshxPath, 0755)
	err = archive.Unzip(archivePath, alshxPath)
	if err != nil {
		logger.Error("ERROR: Unable to download archive")
		os.Exit(1)
	}

	logger.Log("---")
	if args.Install {
		logger.Log("Alshx utilities installed!")
	}
	if args.Update {
		logger.Log("Alshx utilities updated!")
	}
}
