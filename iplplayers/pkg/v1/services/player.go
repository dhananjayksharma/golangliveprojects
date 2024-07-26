package services

import (
	"context"
	"fmt"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/pkg/util"
	"golangliveprojects/iplplayers/pkg/v1/requests"
	"golangliveprojects/iplplayers/pkg/v1/responses"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// want to cacluate some interest for some amount
// interest =  (P × R × T)/100,
type PlayerService interface {
	List(c *gin.Context) (responses.Response, error)
	ListPlayerMatches(c *gin.Context) (responses.Response, error)
	PlayerDetails(c *gin.Context) (responses.Response, error)
	AddPlayer(c *gin.Context) (responses.Response, error)
}

type playerService struct {
	db queries.PersistentSQLDBStorer
}

func NewPlayerService(dbaccess queries.PersistentSQLDBStorer) PlayerService {
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
	service.db.PlayerListQuery(ctx, &playerData)
	responseData.Data = playerData
	responseData.Message = "Player list"
	responseData.RecordSet = nil
	return responseData, nil
}

func (service playerService) PlayerDetails(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	playerCode := c.Param("player_code")
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playerData []responses.PlayerResponse
	service.db.PlayerListQueryPlayerDetails(ctx, &playerData, playerCode)
	responseData.Data = playerData
	responseData.Message = "PlayerDetails"
	responseData.RecordSet = nil
	return responseData, nil
}

func (service playerService) ListPlayerMatches(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	fmt.Println("FINAL LIST PLAYER Here")
	// service.db.List
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playerData []responses.PlayerResponse
	service.db.PlayerListQuery(ctx, &playerData)
	responseData.Data = playerData
	responseData.Message = "ListMatches"
	responseData.RecordSet = nil
	return responseData, nil
}
func (service playerService) AddPlayer(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	var addPlayerRequest requests.PlayerAddRequest
	if err := c.BindJSON(&addPlayerRequest); err != nil {
		log.Println("requests.addPlayerRequest : ", err.Error())
		return responseData, &util.BadRequest{ErrMessage: err.Error()}
	}

	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	// var playerData []responses.PlayerResponse
	// service.db.PlayerListQuery(ctx, &playerData)
	responseData.Data = addPlayerRequest
	responseData.Message = "Player added successfully"
	responseData.RecordSet = nil
	return responseData, nil
}
