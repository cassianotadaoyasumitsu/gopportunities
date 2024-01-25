package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopportunities/docs"
	"gopportunities/handler"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize handler
	handler.InitializerHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		// Show all openings
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
