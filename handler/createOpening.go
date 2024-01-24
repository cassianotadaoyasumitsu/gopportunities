package handler

import "github.com/gin-gonic/gin"

func CreateOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST Opening",
	})
}
