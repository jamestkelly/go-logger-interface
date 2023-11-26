package example

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamestkelly/go-logger-interface"
)

var middlewareLogger = logger.LoggerInterface{Prefix: "MiddlewareService"}

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
