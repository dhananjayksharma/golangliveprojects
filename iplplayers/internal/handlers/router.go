package handlers

import (
	"golangliveprojects/iplplayers/internal/handlers/health"
	"golangliveprojects/iplplayers/internal/handlers/players"
	"golangliveprojects/iplplayers/pkg/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(dbAccess *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/health", health.HealthCheck)

	playerService := services.NewPlayerService(dbAccess)
	playerHandler := players.NewPlayerHandler(playerService)

	router.GET("/players", playerHandler.List)
	router.GET("/players/:id", playerHandler.PlayerDetails)
	router.GET("/players/:id/matches", playerHandler.ListMatches)
	return router
}
