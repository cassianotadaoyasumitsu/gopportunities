package openings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// ShowOpeningHandler
// @Summary Show a new opening
// @Description Show a job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
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

	handler.SendSuccess(ctx, "ShowOpeningHandler", opening)
}
