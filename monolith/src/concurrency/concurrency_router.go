package concurrency

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"runtime"
)

type SearchRequest struct {
	Words []string `json:"words"`
}

func ConcurrencyRouter(router *gin.RouterGroup) {
	g := router.Group("/concurrency")

	engine := NewSearchEngine()

	g.POST("/search", func(c *gin.Context) {
		var req SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fileContents := make(map[string]string)

		_, filename, _, _ := runtime.Caller(0)
		dirPath := filepath.Join(filepath.Dir(filename), "data")
		dir, err := os.Open(dirPath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		files, err := dir.Readdir(0)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		for _, file := range files {
			name := file.Name()
			bytes, err := os.ReadFile(filepath.Join(dirPath, name))
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			fileContents[name] = string(bytes)
		}

		infos, err := engine.SearchWords(req.Words, fileContents)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, infos)
	})
}
