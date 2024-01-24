package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router
	r := gin.Default()
	// Initialize the routes
	initializeRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
