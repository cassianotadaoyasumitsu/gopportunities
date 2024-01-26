package openings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
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
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find the opening
	if err := handler.DB.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Error while finding opening with id: %s", id))
		return
	}
	// Delete the opening
	if err := handler.DB.Delete(&opening).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error while deleting opening with id: %s", id))
		return
	}
	handler.SendSuccess(ctx, "DeleteOpeingHandler", opening)
}
