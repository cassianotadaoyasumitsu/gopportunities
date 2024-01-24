package handler

import "github.com/gin-gonic/gin"

func DeleteOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE Opening",
	})
}
