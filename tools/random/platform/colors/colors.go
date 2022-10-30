package colors

import (
	"fmt"

	"github.com/fatih/color"
)

var Brand = color.New(color.FgMagenta, color.Bold).SprintFunc()
var Error = color.New(color.FgRed, color.Bold).SprintFunc()
var Good = color.New(color.FgGreen, color.Bold).SprintFunc()
var Info = color.New(color.FgBlue, color.Bold).SprintFunc()
var Log = color.New(color.FgWhite).SprintFunc()
var LogBold = color.New(color.FgWhite, color.Bold).SprintFunc()
var WhiteBold = color.New(color.FgWhite, color.Bold).SprintFunc()

func PrintError(messages ...string) {
	if len(messages) == 0 {
		fmt.Println(Error("ERROR: Something went wrong"))
		return
	}
	fmt.Println(Error("ERROR: " + messages[0]))
	if len(messages) == 1 {
		return
	}
	for _, msg := range messages[1:] {
		fmt.Println(Error("      " + msg))
	}
}

func PrintInfo(messages ...string) {
	if len(messages) == 0 {
		return
	}
	fmt.Println(Info("Info: " + messages[0]))
	if len(messages) == 1 {
		return
	}
	for _, msg := range messages[1:] {
		fmt.Println(Log("      " + msg))
	}
}

func PrintGood(messages ...string) {
	if len(messages) == 0 {
		return
	}
	fmt.Println(Good(messages[0]))
	if len(messages) == 1 {
		return
	}
	for _, msg := range messages[1:] {
		fmt.Println(Good(msg))
	}
}

func PrintLog(messages ...string) {
	if len(messages) == 0 {
		return
	}
	fmt.Println(Log(messages[0]))
	if len(messages) == 1 {
		return
	}
	for _, msg := range messages[1:] {
		fmt.Println(Log(msg))
	}
}
