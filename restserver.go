package main

import (
	"github.com/gin-gonic/gin"

	apiv1 "github.com/redru/restserver/api/v1"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Simple group: v1
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/login", apiv1.LoginController)
			// v1.POST("/submit", submitEndpoint)
			// v1.POST("/read", readEndpoint)
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/login", apiv1.LoginController)
			// v2.POST("/submit", submitEndpoint)
			//v2.POST("/read", readEndpoint)
		}
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
