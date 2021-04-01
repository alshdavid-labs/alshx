package scripts

import "os/exec"

func hasCommand(cmd string) bool {
	path, err := exec.LookPath(cmd)
	if err != nil || path == "" {
		return false
	}
	return true
}
