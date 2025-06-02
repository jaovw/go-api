package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jaovw/go-api/handler"

	docs "github.com/jaovw/go-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// [joaovictor - 26/05/2025] Initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		v1.GET("/opening", handler.ListOpeningHandler)

		v1.GET("/opening/:id", handler.ShowOpeningHandler)
	}
	// [joaovictor - 02/06/2025] Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
