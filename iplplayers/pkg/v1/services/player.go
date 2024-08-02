package services

import (
	"context"
	"fmt"
	"golangliveprojects/iplplayers/internal/entities"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/pkg/constants"
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
	UpdatePlayer(c *gin.Context) (responses.Response, error)
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

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	addPlayer := entities.Players{
		PlayerName:     addPlayerRequest.PlayerName,
		Status:         1,
		PlayerCode:     util.GetNewPlayerCode(),
		PlayerCountry:  addPlayerRequest.PlayerCountry,
		PlayerDob:      addPlayerRequest.PlayerDob,
		PlayerCategory: addPlayerRequest.PlayerCategory,
	}

	err := service.db.AddPlayerQuery(ctx, &addPlayer)
	if err != nil {
		return responseData, err
	}
	// playerCode, err :=
	var newPlayer []responses.PlayerResponse
	newAccount := append(newPlayer, responses.PlayerResponse{
		PlayerName:     addPlayerRequest.PlayerName,
		Status:         constants.DB_STATUS_ACTIVE,
		PlayerCode:     addPlayer.PlayerCode,
		PlayerCountry:  addPlayerRequest.PlayerCountry,
		PlayerDob:      addPlayerRequest.PlayerDob,
		PlayerCategory: addPlayerRequest.PlayerCategory,
	})
	responseData.Data = newAccount
	responseData.Message = "Player added successfully"
	responseData.RecordSet = nil
	return responseData, nil
}

func (service playerService) UpdatePlayer(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	playerCode := c.Param("player_code")
	err := util.ValidatePlayerCode(playerCode)
	if err != nil {
		return responseData, &util.BadRequest{ErrMessage: err.Error()}
	}
	var existingData = entities.Players{}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = service.db.GetPlayerByPlayerCode(ctx, &existingData, playerCode)
	if err != nil {
		return responseData, &util.BadRequest{ErrMessage: err.Error()}
	}

	var updatePlayerRequest requests.PlayerUpdateRequest
	if err := c.BindJSON(&updatePlayerRequest); err != nil {
		log.Println("requests.updatePlayerRequest : ", err.Error())
		return responseData, &util.BadRequest{ErrMessage: err.Error()}
	}

	updatePlayerRequestObj := entities.PlayersUpdate{
		PlayerName:     updatePlayerRequest.PlayerName,
		Status:         updatePlayerRequest.PlayerStatus,
		PlayerCountry:  updatePlayerRequest.PlayerCountry,
		PlayerDob:      updatePlayerRequest.PlayerDob,
		PlayerCategory: updatePlayerRequest.PlayerCategory,
	}

	err = service.db.UpdatePlayerQuery(ctx, &updatePlayerRequestObj, playerCode)
	if err != nil {
		return responseData, err
	}

	var updatePlayer []responses.PlayerResponse
	newAccount := append(updatePlayer, responses.PlayerResponse{
		ID:             existingData.ID,
		PlayerName:     updatePlayerRequest.PlayerName,
		Status:         updatePlayerRequest.PlayerStatus,
		PlayerCode:     playerCode,
		PlayerCountry:  updatePlayerRequest.PlayerCountry,
		PlayerDob:      updatePlayerRequest.PlayerDob,
		PlayerCategory: updatePlayerRequest.PlayerCategory,
		CreatedDt:      existingData.CreatedDt,
		UpdatedDt:      existingData.UpdatedDt,
	})
	responseData.Data = newAccount
	responseData.Message = "Player updated successfully"
	responseData.RecordSet = nil
	return responseData, nil
}
