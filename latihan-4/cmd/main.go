package main

import (
	"fmt"
	"latihan-4/config"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users = []User{}

func main() {
	// setup routing
	router := gin.Default()
	router.Use(Trace())

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

	traceID, _ := ctx.Get("tracer_id")
	timestamps := time.Now().Format("2006/01/02 15:04:05")
	traceIDstring := fmt.Sprintf("%s", traceID)
	ctx.Header("X-TRACE-ID", traceIDstring)

	// proses binding
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, err.Error(), ctx.Request.Method, ctx.Request.URL.Path, traceID)

		log.Println(logMessageIncomingRequest)
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

	traceID, _ := ctx.Get("tracer_id")
	timestamps := time.Now().Format("2006/01/02 15:04:05")
	traceIDstring := fmt.Sprintf("%s", traceID)
	ctx.Header("X-TRACE-ID", traceIDstring)

	userID := ctx.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, err.Error(), ctx.Request.Method, ctx.Request.URL.Path, traceID)

		log.Println(logMessageIncomingRequest)
		return
	}

	// proses binding
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, err.Error(), ctx.Request.Method, ctx.Request.URL.Path, traceID)

		log.Println(logMessageIncomingRequest)
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

	traceID, _ := ctx.Get("tracer_id")
	timestamps := time.Now().Format("2006/01/02 15:04:05")
	traceIDstring := fmt.Sprintf("%s", traceID)
	ctx.Header("X-TRACE-ID", traceIDstring)

	if len(users) == 0 {
		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, "data not found", ctx.Request.Method, ctx.Request.URL.Path, traceID)

		log.Println(logMessageIncomingRequest)

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "data not found",
		})
	}

	// tampilkan hasilnya
	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": 201,
		"message":     "get all success",
		"payload":     users,
	})
}

func deleteUser(ctx *gin.Context) {

	traceID, _ := ctx.Get("tracer_id")
	timestamps := time.Now().Format("2006/01/02 15:04:05")
	traceIDstring := fmt.Sprintf("%s", traceID)
	ctx.Header("X-TRACE-ID", traceIDstring)
	userID := ctx.Param("id")
	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, err.Error(), ctx.Request.Method, ctx.Request.URL.Path, traceID)

		log.Println(logMessageIncomingRequest)

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

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// log.Println("Authrorization: set user id with", userId)
		// // get context
		// myCtx := ctx.Request.Context()
		// // set context
		// myCtx = context.WithValue(myCtx, "USER_ID", userId)
		// // get request with new context
		// req := ctx.Request.WithContext(myCtx)
		// // change the request to new request with new context
		// ctx.Request = req

		tracerID := uuid.New()

		timestamps := time.Now().Format("2006/01/02 15:04:05")
		// Format pesan log sebagai string
		logMessageIncomingRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, "incoming Request", ctx.Request.Method, ctx.Request.URL.Path, tracerID)

		log.Println(logMessageIncomingRequest)

		ctx.Set("tracer_id", tracerID)

		ctx.Next()
		// Format pesan log sebagai string
		logMessageAfterRequest := fmt.Sprintf("%s message=\"%s\" method=\"%s\" uri=\"%s\" trace_id=\"%s\"",
			timestamps, "finish request", ctx.Request.Method, ctx.Request.URL.Path, tracerID)

		log.Println(logMessageAfterRequest)

	}
}
