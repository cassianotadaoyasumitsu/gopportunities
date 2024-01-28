package users

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
	"time"
)

func SignInUserHandler(ctx *gin.Context) {
	// Get email and password from request
	user := SignInUserRequest{}

	err := ctx.BindJSON(&user)
	if err != nil {
		return
	}

	// Look up requested user in DB
	users := schemas.User{}
	if err := handler.DB.First(&users, "email = ?", user.Email).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, "Error while finding user")
		return
	}

	// Compare password from request with password from DB
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password)); err != nil {
		handler.SendError(ctx, http.StatusUnauthorized, "Invalid password")
		return
	}

	user.ID = int(users.ID)

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"alg": "HS256",
		"typ": "JWT",
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "Error while generating JWT")
		return
	}

	// Send JWT to client
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", true, true)
	handler.SendSuccess(ctx, "SignInUserHandler", user.ID)
}
