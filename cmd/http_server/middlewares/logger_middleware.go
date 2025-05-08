package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"io"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(config configs.Config) gin.HandlerFunc {
	return StructuredLogger(logging.NewLogger(config))
}

func StructuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/swagger") {
			c.Next()
		} else {
			bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now()
			path := c.FullPath()
			raw := c.Request.URL.RawQuery
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			err := c.Request.Body.Close()
			if err != nil {
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			c.Writer = bodyLogWriter

			c.Next()

			param := gin.LogFormatterParams{}
			param.TimeStamp = time.Now()
			param.Latency = time.Since(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()
			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path

			headers := map[string]string{}
			headers["x-Client-Version"] = c.Request.Header.Get("x-Client-Version")
			headers["x-Client"] = c.Request.Header.Get("x-Client")
			headers["x-App-Version"] = c.Request.Header.Get("x-App-Version")
			headers["language"] = c.Request.Header.Get("language")

			keys := map[logging.ExtraKey]interface{}{}
			keys[logging.Path] = param.Path
			keys[logging.ClientIp] = param.ClientIP
			keys[logging.Method] = param.Method
			keys[logging.Latency] = param.Latency
			keys[logging.StatusCode] = param.StatusCode
			keys[logging.ErrorMessage] = param.ErrorMessage
			keys[logging.BodySize] = param.BodySize
			keys[logging.Headers] = headers

			logger.Info(logging.RequestResponse, logging.API, "Request", keys)
		}
	}
}
