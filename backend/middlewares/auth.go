package middlewares

import (
	"fmt"
	"strings"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/signup" {
			c.Next()
			return;
		}
	
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{
				"error": true,
				"message": "Authorization header is missing",
			})
			return
		}

	
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := []byte(os.Getenv("SECRET"))
		// fmt.Println(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })
	
		if err != nil {
			c.JSON(400, gin.H{
				"error": true,
				"message": "Error decoding token: " + err.Error(),
			})
			return
		}
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
			c.Set("role", claims["role"])
			c.Set("name", claims["name"])
		} else {
			c.JSON(400, gin.H{
				"error": true,
				"message": "Invalid token",
			})
			return
		}
	
		c.Next()
	}
}