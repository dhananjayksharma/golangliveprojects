package handlers

import (
	"golangliveprojects/iplplayers/internal/db/mysql"
	"golangliveprojects/iplplayers/internal/handlers/health"
	"golangliveprojects/iplplayers/internal/handlers/players"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/pkg/v1/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	dbconn := mysql.InitDB()
	dbAccess := queries.NewPersistentSQLDBStore(dbconn)

	router := gin.Default()
	router.GET("/health", health.HealthCheck)

	playerService := services.NewPlayerService(dbAccess)
	playerHandler := players.NewPlayerHandler(playerService)

	router.GET("/players", playerHandler.List)
	router.GET("/players/:player_code", playerHandler.PlayerDetails)
	router.GET("/players/:player_code/matches", playerHandler.ListMatches)
	return router
}
