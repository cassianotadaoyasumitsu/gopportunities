package users

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
)

func SignUpUserHandler(ctx *gin.Context) {
	request := SignUpUserRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Error while validating signup request: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		handler.Logger.Errorf("Error while hashing password: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error while hashing password")
		return
	}

	user := schemas.User{
		Email:    request.Email,
		Password: string(hash),
	}

	if err := handler.DB.Create(&user).Error; err != nil {
		handler.Logger.Errorf("Error while creating a new user(already taken): %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error while creating a new user(already taken)")
		return
	}

	handler.SendSuccess(ctx, "SignupUserHandler", user)
}
