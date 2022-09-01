package Server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cal1co/jpi/algo"
	"github.com/gin-gonic/gin"
)

func JPI(trie *algo.Trie) {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/get/:word", func(c *gin.Context) {
		word := c.Param("word")
		start := time.Now()
		output := algo.Find(word, trie, 3, 0, 0)
		end := time.Since(start).Seconds()
		msg := "results in " + fmt.Sprintf("%f", end) + "s"
		c.JSON(http.StatusOK, gin.H{
			msg: output,
		})
	})
	server.Run()
}
