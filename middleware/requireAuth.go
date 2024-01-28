package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
	"time"
)

func RequireAuth(ctx *gin.Context) {
	// Get the cookie
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized by cookie",
		})
		return
	}

	// Validate the cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})

	claims := token.Claims.(jwt.MapClaims)

	if claims != nil {
		// Check expiration
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized by expiration",
			})
			return
		}
		// Check if the user exists
		user := schemas.User{}
		if err := handler.DB.First(&user, "id = ?", claims["sub"]).Error; err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized by user",
			})
			return
		}

		// Attach the user to the context
		ctx.Set("user", user)

		// Continue
		ctx.Next()
	}
}
