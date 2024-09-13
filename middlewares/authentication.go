package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"main/functions"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	fmt.Println("Authentication")

	return func(c *gin.Context) {
		if len(c.Request.Header["Authorization"]) < 1 {
			log.Println("Error")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		tokenString := c.Request.Header["Authorization"][0]

		if tokenString == "" {
			log.Println("Error")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims, err := functions.VerifyToken(tokenString)
		if err != nil {
			log.Println("Error")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
