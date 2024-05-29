package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gym-management/src/lib"
	"gym-management/src/web/gin/v1/utils"
	"io"
	"strconv"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) Body() string {
	return r.body.String()
}

func readRequestBody(c *gin.Context) string {
	buf, _ := io.ReadAll(c.Request.Body)
	reader := io.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

	body := new(bytes.Buffer)
	_, err := body.ReadFrom(reader)
	if err != nil {
		return ""
	}

	return body.String()
}

func registerOurResponseBodyWriter(c *gin.Context) *responseBodyWriter {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	return w
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		session := utils.ExtractSession(c)

		clientIP := c.ClientIP()
		path := c.Request.RequestURI
		method := c.Request.Method
		userAgent := c.Request.UserAgent()

		requestBody := readRequestBody(c)
		responseWriter := registerOurResponseBodyWriter(c)

		lib.Logger().Info("Start "+method+" "+path, map[string]interface{}{
			"clientIP":    clientIP,
			"userAgent":   userAgent,
			"requestBody": requestBody,
		}, session)

		c.Next()

		statusCode := c.Writer.Status()
		responseBody := responseWriter.Body()

		end := time.Now()
		latency := end.Sub(start)
		latencyInMilliseconds := int64(latency) / int64(time.Millisecond)

		if statusCode >= 500 {
			lib.Logger().Error("End "+method+" "+path, map[string]interface{}{
				"status":       statusCode,
				"latency":      strconv.FormatInt(latencyInMilliseconds, 10) + "ms",
				"responseBody": responseBody,
			}, nil, session)
			return
		}

		if statusCode >= 400 {
			lib.Logger().Warn("End "+method+" "+path, map[string]interface{}{
				"status":       statusCode,
				"latency":      strconv.FormatInt(latencyInMilliseconds, 10) + "ms",
				"responseBody": responseBody,
			}, nil, session)
			return
		}

		lib.Logger().Info("End "+method+" "+path, map[string]interface{}{
			"latency": strconv.FormatInt(latencyInMilliseconds, 10) + "ms",
		}, session)
	}
}
