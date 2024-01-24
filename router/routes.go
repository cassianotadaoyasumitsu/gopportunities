package router

import (
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// Show all openings
		v1.POST("/opening", handler.CreateOpeingHandler)
		v1.GET("/opening", handler.ShowOpeingHandler)
		v1.GET("/openings", handler.ListOpeingHandler)
		v1.PUT("/opening", handler.UpdateOpeingHandler)
		v1.DELETE("/opening", handler.DeleteOpeingHandler)

	}
}
