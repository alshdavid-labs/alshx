package main

import (
	"alshx/src/platform/archive"
	"alshx/src/platform/download"
	"alshx/src/platform/files"
	"alshx/src/platform/github"
	"alshx/src/platform/logging"
	"alshx/src/platform/scripts"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var version = "2"
var latestCommitHash = github.LatestCommitHash()
var homePath = getHomePath()
var alshxPath = filepath.Join(homePath, ".alshx")
var alshxBinPath = filepath.Join(alshxPath, "bin")
var alshxTempPath = filepath.Join(alshxPath, "temp")
var archiveName = "alshx.zip"
var commitHashName = "commit-sha.txt"
var archivePath = filepath.Join(alshxPath, archiveName)
var commitHashPath = filepath.Join(alshxPath, commitHashName)
var remoteArchiveURL = "https://github.com/alshdavid/alshx/archive/master.zip"
var filePermissions fs.FileMode = 0755
var scriptPath = ""

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please include script to execute")
		os.Exit(1)
	}

	var cmdArgs, script, scriptArgs = getArgs()
	var logger = logging.NewLogger(cmdArgs.verbose)
	scriptPath = filepath.Join(alshxBinPath, script)

	if script == "version" {
		logger.Log("Version:", version)
		os.Exit(1)
	}

	if latestCommitHash == "" {
		logger.Log("Error: Unable to get latest commit hash")
		os.Exit(1)
	}

	logger.Info("CurrentHash:", latestCommitHash)
	logger.Info("RootDirectory:", alshxPath)
	logger.Info("Script:", script)
	logger.Info("ScriptPath:", scriptPath)
	logger.Info("Script Args:", scriptArgs)

	prep(logger)

	scripts.Exec(logger, scriptPath, script, scriptArgs)
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

func prep(logger *logging.Logger) {
	if files.NotExists(alshxPath) {
		os.Mkdir(alshxPath, filePermissions)
	}

	if files.Exists(commitHashPath) &&
		files.ReadTextFile(commitHashPath) == latestCommitHash {
		logger.Info("Commit Hash Match: Skipping")
		return
	}

	logger.Log("Commit Hash Different: Downloading")
	if files.Exists(alshxBinPath) {
		os.RemoveAll(alshxBinPath)
	}
	download.DownloadFile(archivePath, remoteArchiveURL)
	archive.Unzip(archivePath, alshxTempPath)
	os.Rename(filepath.Join(alshxTempPath, "alshx-master", "scripts"), alshxBinPath)
	os.WriteFile(commitHashPath, []byte(latestCommitHash), filePermissions)

	os.RemoveAll(archivePath)
	os.RemoveAll(alshxTempPath)
}

type args struct {
	verbose bool
}

func getArgs() (rootArgs args, script string, scriptArgs []string) {

	args := os.Args[1:len(os.Args)]

	targetArgs := false
	for _, arg := range args {
		if !targetArgs && strings.HasPrefix(arg, "-") {
			if arg == "-v" {
				rootArgs.verbose = true
			}
			continue
		} else {
			targetArgs = true
		}
		scriptArgs = append(scriptArgs, arg)
	}

	script = scriptArgs[0]
	if len(scriptArgs) == 1 {
		scriptArgs = []string{}
	} else {
		scriptArgs = scriptArgs[1:]
	}

	return rootArgs, script, scriptArgs
}
