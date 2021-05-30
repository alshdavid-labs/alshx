package handler_help

import (
	"random/src/platform/colors"
)

func Handler() {
	colors.PrintLog(
		"Generate Random: String / Number / GUID",
		"",
		"String:",
		"  random string",
		"  random string --length 20",
		"  random string -l 20",
		"",
		"Number:",
		"  random number",
		"  random number --length 20",
		"  random number -l 20",
		"",
		"GUID:",
		"  random guid",
	)
}
