package main

import (
	"context"
	"latihan-4/config"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []User{}

func main() {
	// setup routing
	router := gin.Default()
	router.Use(Authorization())

	// handler with basic routing
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"version": "not set",
		})
	})

	// api version
	v1 := router.Group("api/v1")
	users := v1.Group("users")
	{
		users.GET("", listUsers)
		users.POST("", addUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}

	router.Run(config.APP_PORT)
}

func addUser(ctx *gin.Context) {
	// variable penampung request
	var req ReqUser

	// proses binding
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := User{
		ID:      int64(rand.Intn(1000)),
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	users = append(users, user)

	// tampilkan hasilnya
	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": 201,
		"message":     "created success",
	})
}

func updateUser(ctx *gin.Context) {
	// variable penampung request
	var req ReqUser

	userID := ctx.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// proses binding
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUsers := []User{}

	for _, user := range users {

		if user.ID == int64(userIDInt) {
			newUser := User{}
			newUser.ID = user.ID
			newUser.Name = req.Name
			newUser.Email = req.Email
			newUser.Address = req.Address

			newUsers = append(newUsers, newUser)
		} else {
			newUsers = append(newUsers, user)

		}

	}

	users = newUsers

	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": 201,
		"message":     "update success",
	})
}

func listUsers(ctx *gin.Context) {

	// tampilkan hasilnya
	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": 201,
		"message":     "get all success",
		"payload":     users,
	})
}

func deleteUser(ctx *gin.Context) {

	userID := ctx.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUsers := []User{}

	for _, user := range users {

		if user.ID != int64(userIDInt) {

			newUsers = append(newUsers, user)
		}
	}

	users = newUsers

	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": 201,
		"message":     "delete success",
	})
}

type ReqUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// do process validation
		// get user id from token
		userId := 10

		log.Println("Authrorization: set user id with", userId)
		// get context
		myCtx := ctx.Request.Context()
		// set context
		myCtx = context.WithValue(myCtx, "USER_ID", userId)
		// get request with new context
		req := ctx.Request.WithContext(myCtx)
		// change the request to new request with new context
		ctx.Request = req
		ctx.Next()
		log.Println("after request")

	}
}
