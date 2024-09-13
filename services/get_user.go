package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	fmt.Println("GetUser")

	claims, _ := c.Get("claims")

	c.JSON(http.StatusOK, gin.H{
		"user":   "user",
		"claims": claims,
	})
}
