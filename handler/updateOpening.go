package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error while validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Error while finding opening with id: %s", id))
		return
	}

	// Update the opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save the opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error while updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error while updating opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "UpdateOpeningHandler", opening)
}
