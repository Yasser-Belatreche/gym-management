package scratch

import (
	"github.com/gin-gonic/gin"
	"os"
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

		dir, err := os.Open("data")
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
			fileName := file.Name()
			filePath := "data/" + fileName
			bytes, err := os.ReadFile(filePath)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			fileContents[fileName] = string(bytes)
		}

		infos, err := engine.SearchWords(req.Words, fileContents)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, infos)
	})
}
