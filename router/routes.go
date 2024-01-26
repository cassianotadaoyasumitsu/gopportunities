package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopportunities/docs"
	"gopportunities/handler"
	"gopportunities/handler/openings"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize handler
	handler.InitializerHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		// Show all openings
		v1.POST("/opening", openings.CreateOpeningHandler)
		v1.GET("/opening", openings.ShowOpeningHandler)
		v1.GET("/openings", openings.ListOpeningHandler)
		v1.PUT("/opening", openings.UpdateOpeningHandler)
		v1.DELETE("/opening", openings.DeleteOpeningHandler)
	}
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
