package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(handlers *Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// Use /api/v1/produce as the route group for API versioning
	produceGroup := router.Group("/api/v1/produce")
	{
		produceGroup.POST("/", handlers.AddProduce)
		produceGroup.GET("/:code", handlers.GetProduceByCode)
		produceGroup.GET("/", handlers.SearchProduce)
		produceGroup.DELETE("/:code", handlers.DeleteProduce)
	}

	return router
}
