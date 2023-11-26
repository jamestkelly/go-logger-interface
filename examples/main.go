package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamestkelly/go-logger-interface"
)

var (
	middlewareLogger = logger.LoggerInterface{Prefix: "MiddlewareService"}
	exampleLogger    = logger.LoggerInterface{Prefix: "ExampleLogger"}
)

// LogRequest
// Custom example middleware logger for logging of requests made to a Gin API server.
func LogRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()
		lag := time.Since(t)

		middlewareLogger.LogMessage(
			fmt.Sprintf("%s %s %s %s", ctx.Request.Method, ctx.Request.RequestURI, ctx.Request.Proto, lag),
			"INFO",
		)
	}
}

// main
// Example main function to log a bunch of statements using LoggerInterface
func main() {
	exampleLogger.LogMessage(
		"This is a 'DEBUG' message.",
		"DEBUG",
	)

	exampleLogger.LogMessage(
		"This is an 'INFO' message.",
		"INFO",
	)

	exampleLogger.LogMessage(
		fmt.Sprintf("This is a %s 'INFO' message.", "formatted"),
		"INFO",
	)

	exampleLogger.LogMessage(
		"This is a 'WARN' message.",
		"WARN",
	)

	exampleLogger.LogMessage(
		"This is an 'ERROR' message.",
		"ERROR",
	)

	exampleLogger.LogMessage(
		"This is a 'FATAL' message.",
		"FATAL",
	)
}
