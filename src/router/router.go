package router

import (
	"etl-neoway-challenge/src/handlers"
	"etl-neoway-challenge/src/middleware"

	"github.com/gin-gonic/gin"
)

//StartRouter - Inicializa router e os handlers
func StartRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	var publicGroup = router.Group("/etl-golang/v1")

	publicGroup.POST("/file", handlers.ExtractTransformLoad)

	return router
}
