// version_prod.go
// +build prod

package main

import (
	_ "embed"
)

//go:embed version.txt
var version string
