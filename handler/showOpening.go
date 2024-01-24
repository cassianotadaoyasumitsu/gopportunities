package handler

import "github.com/gin-gonic/gin"

func ShowOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Opening",
	})
}
