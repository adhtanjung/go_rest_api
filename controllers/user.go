package controllers

import (
	"fmt"
	"rest_api_gin/dto"
	"time"

	"context"
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DATABASE INSTANCE

var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

// ShowAccount godoc
// @Summary Show a account
// @Description get all users
// @Produce  json
// @Success 200 {object} []User
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, &dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error while getting all users",
		})
		return
	}

	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, &dto.ResponseWithData{
		Status:  http.StatusOK,
		Message: "Successfully got all users",
		Data:    users,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User

	c.BindJSON(&user)
	fmt.Println(user)

	username := user.Username
	email := user.Email
	password := user.Password

	id := guuid.New().String()

	newUser := User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error while creating user, reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error while creating user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created user",
	})
	return

}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} User
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /user/{userId} [get]
func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")
	user := User{}

	err := collection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)

	if err != nil {
		log.Printf("Error while getting single user, reason: %v\n", err)
		c.JSON(http.StatusNotFound, &dto.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully got single user",
		"data":    user,
	})
	return

}

func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	var user User

	c.BindJSON(&user)

	username := user.Username
	newData := bson.M{"$set": bson.M{"username": username, "updated_at": time.Now()}}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": userId}, newData)

	if err != nil {
		log.Printf("Error while editing user, reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error while editing user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully edited user",
	})

	return
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": userId})

	if err != nil {
		log.Printf("Error while deleting user, reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error while deleting user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully deleted user",
	})
	return
}
