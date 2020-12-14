package handlers

import (
	"etl-neoway-challenge/src/services"

	"github.com/gin-gonic/gin"
)

func ExtractTransformLoad(c *gin.Context) {
	services.Etl()
}
