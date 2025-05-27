package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	// [joaovictor - 26/05/2025] Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/opening", handler.ListOpeningHandler)
	}
}
