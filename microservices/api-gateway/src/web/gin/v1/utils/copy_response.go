package utils

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CopyResponse(res *http.Response, c *gin.Context) {
	for k, v := range res.Header {
		c.Header(k, v[0])
	}

	c.Status(res.StatusCode)
	io.Copy(c.Writer, res.Body)
}
