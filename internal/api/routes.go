package api

import "github.com/gin-gonic/gin"

func SetupRoutes(handlers *Handlers) *gin.Engine {
	router := gin.Default()

	produceGroup := router.Group("/produce")
	{
		produceGroup.POST("/", handlers.AddProduce)
		produceGroup.GET("/:code", handlers.GetProduceByCode)
		produceGroup.GET("/", handlers.SearchProduce)
		produceGroup.DELETE("/:code", handlers.DeleteProduce)
	}

	return router
}
