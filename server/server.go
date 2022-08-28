package Server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JPI() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/get/:word", func(c *gin.Context) {
		word := c.Param("word")
		c.JSON(http.StatusOK, gin.H{
			"results": word,
		})
	})
	server.Run()
}
