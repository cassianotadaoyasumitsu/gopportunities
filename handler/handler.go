package handler

import "github.com/gin-gonic/gin"

func CreateOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST Opening",
	})
}

func ShowOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Opening",
	})
}

func ListOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET Openings",
	})
}

func UpdateOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT Opening",
	})
}

func DeleteOpeingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE Opening",
	})
}
