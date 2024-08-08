package handlers

import (
	"golangliveprojects/iplplayers/internal/db/mysql"
	"golangliveprojects/iplplayers/internal/db/rediscachedb"
	"golangliveprojects/iplplayers/internal/handlers/health"
	"golangliveprojects/iplplayers/internal/handlers/players"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/internal/queriesredis"
	"golangliveprojects/iplplayers/pkg/v1/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	dbconn := mysql.InitDB()
	dbAccess := queries.NewPersistentSQLDBStore(dbconn)

	redisConn := rediscachedb.InitDBRedis()
	redisDBAccess := queriesredis.NewRedisCacheDBStore(redisConn)

	router := gin.Default()
	router.GET("/health", health.HealthCheck)

	playerService := services.NewPlayerService(dbAccess, redisDBAccess)
	playerHandler := players.NewPlayerHandler(playerService)

	v1PlayerGroup := router.Group("/v1/players")
	{
		v1PlayerGroup.GET("/", playerHandler.List)
		v1PlayerGroup.GET("/:player_code", playerHandler.PlayerDetails)

		v1PlayerGroup.GET("/:player_code/matches", playerHandler.ListPlayerMatches)
		v1PlayerGroup.POST("/", playerHandler.AddPlayer)
		v1PlayerGroup.PUT("/:player_code", playerHandler.UpdatePlayer)

		v1PlayerGroup.GET("/registration/activation/:player_code/:activate_code", playerHandler.PlayerActivate)
	}

	v1MatcheGroup := router.Group("/v1/matches")
	{
		v1MatcheGroup.GET("/", playerHandler.List)
		v1MatcheGroup.GET("/:player_code", playerHandler.PlayerDetails)
		v1MatcheGroup.GET("/:player_code/matches", playerHandler.ListPlayerMatches)
		v1MatcheGroup.POST("/", playerHandler.AddPlayer)
	}

	return router
}
