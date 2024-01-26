package openings

import (
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
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

	if err := handler.DB.Find(&openings).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "Error while finding openings")
		return
	}

	handler.SendSuccess(ctx, "ListOpeningHandler", openings)
}
