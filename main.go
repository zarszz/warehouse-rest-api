package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/utils"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "hello world",
		})
	})
	router.Run(utils.GetEnvVariable("LISTEN_PORT"))
}
