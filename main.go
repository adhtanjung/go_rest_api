package main

import (
	"log"
	"rest_api_gin/config"
	"rest_api_gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	// router.Run(":8080")

	// route handlers / endpoints
	routes.Routes(router)

	// start the server
	log.Fatal(router.Run(":8080"))

}
