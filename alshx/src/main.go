package main

import (
	"alshx/src/platform/archive"
	"alshx/src/platform/download"
	"alshx/src/platform/files"
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

// var latestCommitHash = github.LatestCommitHash()
var homePath = getHomePath()
var alshxPath = filepath.Join(homePath, ".alshx")
var alshxBinPath = filepath.Join(alshxPath, "bin")
var alshxBinCachePath = filepath.Join(alshxPath, "bin-cache")
var alshxTempPath = filepath.Join(alshxPath, "temp")
var archiveName = "alshx.zip"
var commitHashName = "commit-sha.txt"
var archivePath = filepath.Join(alshxPath, archiveName)
var commitHashPath = filepath.Join(alshxPath, commitHashName)
var remoteArchiveURL = "https://github.com/alshdavid/alshx/archive/master.zip"
var filePermissions fs.FileMode = 0755
var scriptPath = ""
var scriptCachePath = ""

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please include script to execute")
		os.Exit(1)
	}

	var cmdArgs, script, scriptArgs = getArgs()
	var logger = logging.NewLogger(cmdArgs.verbose)
	scriptPath = filepath.Join(alshxBinPath, script)
	scriptCachePath = filepath.Join(alshxBinCachePath, script)

	if script == "version" {
		logger.Log("Version:", version)
		os.Exit(1)
	}

	if script == "clean" {
		logger.Log("Cleaning:", alshxPath)
		os.RemoveAll(alshxPath)
		os.Exit(1)
	}

	if script == "update" {
		logger.Log("Updating stash:", alshxPath)
		updateStash(logger)
		os.Exit(1)
	}

	logger.Info("RootDirectory:", alshxPath)
	logger.Info("Script:", script)
	logger.Info("ScriptPath:", scriptPath)
	logger.Info("Script Args:", scriptArgs)

	if files.NotExists(alshxPath) {
		updateStash(logger)
	}

	scripts.Exec(logger, scriptPath, scriptCachePath, script, scriptArgs)
}

func updateStash(logger *logging.Logger) {
	if files.NotExists(alshxPath) {
		os.Mkdir(alshxPath, filePermissions)
	}

	if files.Exists(alshxBinCachePath) {
		logger.Info("INFO: Deleting Bin Cache")
		os.RemoveAll(alshxBinCachePath)
	}

	logger.Info("INFO: Downloading repo")
	err := download.DownloadFile(archivePath, remoteArchiveURL)
	if err != nil {
		logger.Log("Failed to download archive", err)
		os.Exit(1)
	}
	archive.Unzip(archivePath, alshxTempPath)
	logger.Info("INFO: Creating bin cache folder")
	scriptsPath := filepath.Join(alshxTempPath, "alshx-master", "scripts")
	files.CopyDir(scriptsPath, alshxBinCachePath)

	if files.NotExists(alshxBinPath) {
		logger.Info("INFO: Creating bin folder")
		files.CopyDir(scriptsPath, alshxBinPath)
	}

	logger.Info("INFO: Cleaning up")
	os.RemoveAll(alshxTempPath)
	os.RemoveAll(archivePath)
	os.RemoveAll(commitHashPath)
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
