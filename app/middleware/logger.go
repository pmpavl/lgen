package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		logRequest := log.Debug().
			Str("method", c.Request.Method).
			Str("uri", c.Request.URL.Path).
			Str("metadata.UserAgent", c.Request.UserAgent())

		// Before request
		c.Next()
		// After request

		logRequest.
			Float64("latency_ms", float64(time.Since(startTime).Nanoseconds())/float64(time.Millisecond)).
			Int("status", c.Writer.Status()).
			Msg("request")
	}
}
