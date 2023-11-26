package logger

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/fatih/color"
)

// TestFormatSeverity
// Unit test to verify that the formatSeverity method correctly formats strings with the given
// colour.
func TestFormatSeverity(t *testing.T) {
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan = color.New(color.FgCyan).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red = color.New(color.FgRed).SprintFunc()
	boldRed = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()

	testStrings := []struct {
		text     string
		expected string
	}{
		{text: "DEBUG", expected: magenta("DEBUG")},
		{text: "INFO", expected: cyan("INFO")},
		{text: "WARN", expected: yellow("WARN")},
		{text: "ERROR", expected: red("ERROR")},
		{text: "FATAL", expected: boldRed("FATAL")},
	}

	for _, s := range testStrings {
		result := formatSeverity(s.text)

		if result != s.expected {
			t.Errorf(
				"Expected formatted severity for Severity(%s) to return '%s', got '%s' instead.",
				s.text,
				s.expected,
				result,
			)
		}
	}
}

// TestIsNullOrEmptyString
// Unit test to verify that the IsNullOrEmptyString method correctly identifies null or empty strings.
func TestIsNullOrEmptyString(t *testing.T) {
	str := ""
	if !isNullOrEmptyString(str) {
		t.Errorf("Expected input string '%s' to result in 'true', got '%v' instead.", str, isNullOrEmptyString(str))
	}

	str = "Hello world!"
	if isNullOrEmptyString(str) {
		t.Errorf("Expected input string '%s' to result in 'false', got '%v' instead.", str, isNullOrEmptyString(str))
	}
}

// MockLoggerInterface is a mock implementation of LoggerInterface for testing purposes.
type MockLoggerInterface struct {
	Prefix string
}

// MockLogMessageBuffer
// Helper function to capture the output of LogMessage.
func (l *MockLoggerInterface) MockLogMessageBuffer(message, severity string) *bytes.Buffer {
	var buf bytes.Buffer
	l.MockLogMessage(&buf, message, severity)
	return &buf
}

// MockLogMessage
// Modified version of the original LogMessage method to write to a specified buffer.
func (l *MockLoggerInterface) MockLogMessage(output *bytes.Buffer, message, severity string) {
	if isNullOrEmptyString(message) {
		l.MockLogMessage(output, "No message string passed to logger interface.", "ERROR")
	}

	if isNullOrEmptyString(severity) {
		l.MockLogMessage(output, "No severity string passed to logger interface.", "ERROR")
	}

	logTime := time.Now().Format("2006-01-02 15:04:05")

	if severity == "FATAL" {
		fmt.Fprintf(output, base, logTime, formatSeverity(severity), l.Prefix, cyan(message))
		return
	}

	fmt.Fprintf(output, base, logTime, formatSeverity(severity), l.Prefix, cyan(message))
}

// TestLogMessage
// Unit test to verify through the implementation of mock functionality that the `LogMessage`
// method correctly prints logs to the standard output.
func TestLogMessage(t *testing.T) {
	// Create a MockLoggerInterface instance
	mockLogger := &MockLoggerInterface{
		Prefix: "TestPrefix",
	}

	// Test case 1: Log a regular message
	message := "Test Message"
	severity := "INFO"
	expectedOutput := fmt.Sprintf(base, time.Now().Format("2006-01-02 15:04:05"), formatSeverity(severity), mockLogger.Prefix, cyan(message))

	outputBuffer := mockLogger.MockLogMessageBuffer(message, severity)
	actualOutput := outputBuffer.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s\nActual output: %s", expectedOutput, actualOutput)
	}

	// Test case 2: Log a FATAL message
	message = "Fatal Error"
	severity = "FATAL"
	expectedOutput = fmt.Sprintf(base, time.Now().Format("2006-01-02 15:04:05"), formatSeverity(severity), mockLogger.Prefix, cyan(message))

	outputBuffer = mockLogger.MockLogMessageBuffer(message, severity)
	actualOutput = outputBuffer.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s\nActual output: %s", expectedOutput, actualOutput)
	}
}
