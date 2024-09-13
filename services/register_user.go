package services

import (
	"fmt"
	"log"
	"main/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	fmt.Println("RegisterUser")
	log.Println("RegisterUser")

	inputs := map[string]string{}

	if err := c.ShouldBindJSON(&inputs); err != nil {
		log.Println("Error ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	inputs["hash"], _ = functions.HashPassword(inputs["password"])

	c.JSON(http.StatusOK, gin.H{
		"message": inputs,
	})
}
