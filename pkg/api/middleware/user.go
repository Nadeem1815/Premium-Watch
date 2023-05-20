package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("UserAuth")
	// Todo:check if user if block in database

	if err != nil {

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID, err := ValidateToken(tokenString)
	if err != nil {

		fmt.Println("user id", userID)
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	c.Set("userID", userID)
	c.Next()
}
