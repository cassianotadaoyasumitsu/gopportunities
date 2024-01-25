package handler

import (
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error while validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error while creating a new opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error while creating a new opening")
		return
	}

	sendSuccess(ctx, "CreateOpeningHandler", opening)
}
