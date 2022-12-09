package Server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	cache "github.com/cal1co/jpi/cache"
	controller "github.com/cal1co/jpi/controllers"
	"github.com/gin-gonic/gin"
)

func JPI(trieData controller.Output) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	router.GET("/get/:word", func(c *gin.Context) {
		word := strings.ToLower(c.Param("word"))
		start := time.Now()
		res := cache.LookupAndCache(trieData, word)
		end := time.Since(start).Seconds()

		msg := strconv.Itoa(len(res)) + " results for " + "'" + word + "'" + " in " + fmt.Sprintf("%f", end) + "s"

		c.JSON(http.StatusOK, gin.H{
			msg: res,
		})
	})
	router.Run()
}
