package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const nsToms float64 = 1000 * 1000

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
			Float64("latency_ms", float64(time.Since(startTime).Nanoseconds())/nsToms).
			Int("status", c.Writer.Status()).
			Msg("request")
	}
}
