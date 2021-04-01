package files

import "os"

func ReadTextFile(path string) string {
	file, _ := os.ReadFile(path)
	return string(file)
}
