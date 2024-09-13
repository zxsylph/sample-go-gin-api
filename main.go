package main

import (
	"fmt"
	"net/http"

	"main/middlewares"
	"main/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	basePath := "/"

	server := gin.Default()

	root := server.Group(basePath)
	authorized := root.Group("/")
	authorized.Use(middlewares.Authentication())

	root.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	root.POST("/upload", services.Upload)
	root.POST("/registeruser", services.RegisterUser)
	root.POST("/login", services.Login)

	authorized.GET("/user", services.GetUser)

	server.Run()
}
