package openings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// UpdateOpeningHandler
// @Summary Update a new opening
// @Description Update a job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	_, userID := handler.GetUserInfo(ctx)
	request := UpdateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Error while validating request: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := handler.DB.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Error while finding opening with id: %s", id))
		return
	}

	if opening.UserID != userID {
		handler.SendError(ctx, http.StatusForbidden, "You are not allowed to update this opening")
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
	if err := handler.DB.Save(&opening).Error; err != nil {
		handler.Logger.Errorf("Error while updating opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error while updating opening with id: %s", id))
		return
	}

	handler.SendSuccess(ctx, "UpdateOpeningHandler", opening)
}
