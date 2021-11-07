package main

import (
	"go_rest_gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

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
