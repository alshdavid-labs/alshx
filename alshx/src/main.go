package main

import (
	"alshx/src/platform/archive"
	"alshx/src/platform/download"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var homePath = getHomePath()
var alshxPath = filepath.Join(homePath, ".alshx")
var alshxBinPath = filepath.Join(alshxPath, "bin")
var alshxTempPath = filepath.Join(alshxPath, "temp")
var archiveName = "alshx.zip"
var archivePath = filepath.Join(alshxPath, archiveName)
var remoteArchiveURL = "https://github.com/alshdavid/alshx/archive/master.zip"
var script = os.Args[1]
var scriptArgs = os.Args[2:len(os.Args)]

func main() {
	if _, err := os.Stat(alshxPath); os.IsNotExist(err) {
		os.Mkdir(alshxPath, 0755)
		os.Mkdir(alshxTempPath, 0755)
		updateBin()
	}

	fmt.Println(alshxPath)
	fmt.Println(os.Args)
	fmt.Println(script)
	fmt.Println(scriptArgs)

	cmd := exec.Command("echo", "ok")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func getHomePath() string {
	if os.Getenv("ALSHX_PATH") != "" {
		return os.Getenv("ALSHX_PATH")
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

func updateBin() {
	if _, err := os.Stat(alshxBinPath); !os.IsNotExist(err) {
		os.RemoveAll(alshxBinPath)
	}
	download.DownloadFile(archivePath, remoteArchiveURL)
	archive.Unzip(archivePath, alshxTempPath)
	os.Rename(filepath.Join(alshxTempPath, "alshx-master", "scripts"), alshxBinPath)
	os.RemoveAll(alshxTempPath)
}
