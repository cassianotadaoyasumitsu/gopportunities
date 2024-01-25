package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/schemas"
	"net/http"
)

func ShowOpeningHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "ShowOpeningHandler", opening)
}
