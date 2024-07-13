package handlers

import (
	"golangliveprojects/iplplayers/internal/handlers/health"
	"golangliveprojects/iplplayers/internal/handlers/matches"
	"golangliveprojects/iplplayers/internal/handlers/players"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(DB *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/health", health.HealthCheck)
	router.GET("/list", players.List)
	router.GET("/matchlist", matches.List)
	return router
}
