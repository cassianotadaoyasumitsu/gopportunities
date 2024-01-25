package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// DeleteOpeningHandler
// @Summary Delete a new opening
// @Description Delete a job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find the opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Error while finding opening with id: %s", id))
		return
	}
	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error while deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "DeleteOpeingHandler", opening)
}
