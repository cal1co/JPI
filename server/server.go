package Server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	data "github.com/cal1co/jpi/handleData"
	"github.com/gin-gonic/gin"
)

func JPI(trieData data.Output) {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/get/:word", func(c *gin.Context) {
		word := c.Param("word")

		start := time.Now()
		res := data.Fetch(&trieData.Nodes, trieData.Slice, word)
		end := time.Since(start).Seconds()

		msg := strconv.Itoa(len(res)) + " results for " + "'" + word + "'" + " in " + fmt.Sprintf("%f", end) + "s"

		c.JSON(http.StatusOK, gin.H{
			msg: res,
		})
	})
	server.Run()
}
