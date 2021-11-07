package routes

import (
	"net/http"
	"rest_api_gin/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/", welcome)
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userId", controllers.GetSingleUser)
	router.PUT("/user/:userId", controllers.EditUser)
	router.DELETE("/user/:userId", controllers.DeleteUser)
	router.NoRoute(notFound)

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
