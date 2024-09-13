package services

import (
	"fmt"
	"log"
	"net/http"

	"main/functions"

	"github.com/gin-gonic/gin"
)

var validUser = map[string]string{
	"username": "username",
	"hash":     "$2a$14$r2J3S4zdt2WMCzKBgy7r7umyjJcfznKbxVIsKFM5V1O3GXUanUBQ.",
}

func Login(c *gin.Context) {
	fmt.Println("Login")
	log.Println("Login")

	inputs := map[string]string{}

	if err := c.ShouldBindJSON(&inputs); err != nil {
		log.Println("Error ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	if inputs["username"] != validUser["username"] || !functions.CheckPasswordHash(inputs["password"], validUser["hash"]) {
		log.Println("Error Invalid User")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	token, err := functions.CreateToken(int64(1), inputs["username"])

	fmt.Printf("err: %v\n", err)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}
