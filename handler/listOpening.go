package handler

import (
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error while finding openings")
		return
	}

	sendSuccess(ctx, "ListOpeingHandler", openings)
}
