package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func InitMiddlewareLogger() (*os.File, error) {
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile("logs/middleware.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)

	return logFile, nil
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var requestBody []byte
		if c.Request.Method != http.MethodGet {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		w := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[GIN] %s %s | %d | %s | Request: %s | Response: %s\n",
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			duration,
			string(requestBody),
			w.body.String(),
		)
	}
}
