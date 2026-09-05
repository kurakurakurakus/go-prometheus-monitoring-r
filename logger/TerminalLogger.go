package logger

import "fmt"

type TerminalLogger struct{}

func NewTerminalLogger() *TerminalLogger {
	return &TerminalLogger{}
}

func Log(s string) {
	fmt.Printf("[Terminal Logger] %v", s)
}
