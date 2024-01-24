package router

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// Show all openings
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "GET Opening",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "POST Opening",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "DELETE Opening",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "PUT Opening",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "GET Openings",
			})
		})

	}
}
