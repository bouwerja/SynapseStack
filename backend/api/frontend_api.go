package api

import "github.com/gin-gonic/gin"

func RunApi() {
	ba := BasicApi()

	ba.Run(":8080")
}

func BasicApi() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
