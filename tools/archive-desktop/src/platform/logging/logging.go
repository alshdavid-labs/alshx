package logging

import "fmt"

type Logger struct {
	verbose bool
}

func NewLogger(verbose bool) *Logger {
	return &Logger{
		verbose,
	}
}

func (l *Logger) Info(args ...interface{}) {
	if l.verbose {
		fmt.Println(args...)
	}
}

func (l *Logger) Log(args ...interface{}) {
	fmt.Println(args...)
}
