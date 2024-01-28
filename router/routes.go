package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopportunities/docs"
	"gopportunities/handler"
	"gopportunities/handler/openings"
	"gopportunities/handler/users"
	"gopportunities/middleware"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize handler
	handler.InitializerHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Routes
	v1 := r.Group(basePath)
	{
		// Users
		v1.POST("/signup", users.SignUpUserHandler)
		v1.POST("/signing", users.SignInUserHandler)

		// Show all openings
		v1.POST("/opening", middleware.RequireAuth, openings.CreateOpeningHandler)
		v1.GET("/opening", openings.ShowOpeningHandler)
		v1.GET("/openings", openings.ListOpeningHandler)
		v1.PUT("/opening", middleware.RequireAuth, openings.UpdateOpeningHandler)
		v1.DELETE("/opening", middleware.RequireAuth, openings.DeleteOpeningHandler)
	}
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
