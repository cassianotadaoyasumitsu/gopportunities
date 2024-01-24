package handler

import "github.com/gin-gonic/gin"

func UpdateOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT Opening",
	})
}
