package logger

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

// base The base formatted string to be printed by the Logger.
// magenta A magenta colour utilised for "DEBUG" level logs.
// cyan A cyan colour utilised for standard output to the Logger.
// green A green colour utilised for "INFO" level logs.
// yellow A yellow colour utilised for "WARN" level logs.
// red A red colour utilised for "ERROR" level logs.
// boldRed A bold red colour utilised for "FATAL" level logs.
var (
	base    = "[%s] %s [%s] - %s\r\n"
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	boldRed = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
)

// LoggerInterface
// A structure containing the Prefix or name of the service that logs are written for.
type LoggerInterface struct {
	Prefix string
}

// isNullOrEmptyString
// Returns a boolean value indicating if a string is either empty or null.
func isNullOrEmptyString(str string) bool {
	if str == "" || len(str) <= 0 {
		return true
	}

	return false
}

// formatSeverity
// Private method to convert the provided severity string to its corresponding colour for printing.
func formatSeverity(severity string) string {
	var formattedSeverity string

	switch severity {
	case "DEBUG":
		formattedSeverity = magenta(severity)
	case "INFO":
		formattedSeverity = green(severity)
	case "WARN":
		formattedSeverity = yellow(severity)
	case "ERROR":
		formattedSeverity = red(severity)
	case "FATAL":
		formattedSeverity = boldRed(severity)
	}

	return formattedSeverity
}

// LogMessage
// Structure method to print a log message provided both a message and severity.
func (l *LoggerInterface) LogMessage(message, severity string) {
	if isNullOrEmptyString(message) {
		l.LogMessage("No message string passed to logger interface.", "ERROR")
	}

	if isNullOrEmptyString(severity) {
		l.LogMessage("No severity string passed to logger interface.", "ERROR")
	}

	logTime := time.Now().Format("2006-01-02 15:04:05")

	if severity == "FATAL" {
		log.Fatalf(base, logTime, formatSeverity(severity), l.Prefix, cyan(message))
		return
	}

	fmt.Printf(base, logTime, formatSeverity(severity), l.Prefix, cyan(message))
}
