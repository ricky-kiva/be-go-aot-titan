package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	SECRET = "secret"
)

func AuthValidation(c *gin.Context) {
	// get `string token` from Header
	tokenString := c.Request.Header.Get("Authorization")

	// check if token "missing"
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Missing token",
		})
		c.Abort()
		return
	}

	// parse `string token` to check if the "token" is valid
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// verify token signature validity & signing method
		if _, valid := t.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("invalid token: %s", t.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	// response on token validity
	if token != nil && err == nil {
		fmt.Println("Token verified")
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		c.Abort()
	}
}
