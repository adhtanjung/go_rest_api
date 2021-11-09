package routes

import (
	"net/http"
	"rest_api_gin/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/api", welcome)
	router.GET("/api/users", controllers.GetAllUsers)
	router.POST("/api/user", controllers.CreateUser)
	router.GET("/api/ser/:userId", controllers.GetSingleUser)
	router.PUT("/api/user/:userId", controllers.EditUser)
	router.DELETE("/api/user/:userId", controllers.DeleteUser)
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
