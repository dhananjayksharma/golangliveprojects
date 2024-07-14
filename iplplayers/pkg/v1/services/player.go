package services

import (
	"context"
	"fmt"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/pkg/v1/responses"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerService interface {
	List(c *gin.Context) (responses.Response, error)
	ListMatches(c *gin.Context) (responses.Response, error)
	PlayerDetails(c *gin.Context) (responses.Response, error)
}

type playerService struct {
	// db queries.PersistentSQLDBStorer
	db *gorm.DB
}

func NewPlayerService(dbaccess *gorm.DB) PlayerService {
	return &playerService{db: dbaccess}
}

func (service playerService) List(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	fmt.Println("FINAL LIST PLAYER Here")
	// service.db.List
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playerData []responses.PlayerResponse
	queries.ListQuery(ctx, &playerData, service.db)
	responseData.Data = playerData
	responseData.Message = "Player list"
	responseData.RecordSet = nil
	return responseData, nil
}

func (service playerService) PlayerDetails(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	fmt.Println("FINAL LIST PLAYER Here")
	// service.db.List
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playerData []responses.PlayerResponse
	queries.ListQuery(ctx, &playerData, service.db)
	responseData.Data = playerData
	responseData.Message = "PlayerDetails"
	responseData.RecordSet = nil
	return responseData, nil
}

func (service playerService) ListMatches(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	fmt.Println("FINAL LIST PLAYER Here")
	// service.db.List
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playerData []responses.PlayerResponse
	queries.ListQuery(ctx, &playerData, service.db)
	responseData.Data = playerData
	responseData.Message = "ListMatches"
	responseData.RecordSet = nil
	return responseData, nil
}
