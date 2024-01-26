package handler

import (
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// ListOpeningHandler
// @Summary List a new opening
// @Description List a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := dbp.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error while finding openings")
		return
	}

	sendSuccess(ctx, "ListOpeingHandler", openings)
}
