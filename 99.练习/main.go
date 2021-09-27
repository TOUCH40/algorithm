package main

import "fmt"

type logClosure func(format string, v ...interface{})

func LoggerWrapper(logType string) logClosure {
	return func(format string, v ...interface{}) {
		fmt.Printf(fmt.Sprintf("[%s]%s", logType, format), v...)
		fmt.Println()
	}
}

func main() {
	info_logger := LoggerWrapper("INFO")
	warning_logger := LoggerWrapper("WARNING")
	info_logger("this is a %s log", "info")
	warning_logger("this is a %s log", "warning")
}
