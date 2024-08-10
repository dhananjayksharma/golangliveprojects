package services

import (
	"context"
	"fmt"
	"golangliveprojects/iplplayers/internal/entities"
	"golangliveprojects/iplplayers/internal/messages"
	"golangliveprojects/iplplayers/internal/queries"
	"golangliveprojects/iplplayers/internal/queriesredis"
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
	PlayerActivate(c *gin.Context) (responses.Response, error)
}

type playerService struct {
	db            queries.PersistentSQLDBStorer
	redisDBAccess queriesredis.RedisCacheDBStorer
}

func NewPlayerService(dbaccess queries.PersistentSQLDBStorer, redisDBAccess queriesredis.RedisCacheDBStorer) PlayerService {
	return &playerService{db: dbaccess, redisDBAccess: redisDBAccess}
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
		Status:         constants.DB_STATUS_REGISTERED,
		PlayerCode:     addPlayer.PlayerCode,
		PlayerCountry:  addPlayerRequest.PlayerCountry,
		PlayerDob:      addPlayerRequest.PlayerDob,
		PlayerCategory: addPlayerRequest.PlayerCategory,
	})
	activationCode, err := util.GenerateActivationCode()
	if err != nil {
		return responseData, err
	}
	keyName := util.GetPlayerActicationKey(addPlayer.PlayerCode)
	err = service.redisDBAccess.SaveRegistrationDataByKey(ctx, keyName, activationCode)
	if err != nil {
		return responseData, err
	}
	// send email

	messages.SendEmail(activationCode)

	// registrationKey, _ := service.redisDBAccess.GetRegistrationDataByKey(ctx, addPlayer.PlayerCode)

	responseData.Data = newAccount
	responseData.Message = fmt.Sprintf("Player added successfully activation key:%s", activationCode)
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

func (service playerService) PlayerActivate(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response
	playerCode := c.Param("player_code")
	activateCode := c.Param("activate_code")
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
	keyName := util.GetPlayerActicationKey(existingData.PlayerCode)
	code, err := service.redisDBAccess.GetRegistrationDataByKey(ctx, keyName)
	if err != nil {
		return responseData, &util.BadRequest{ErrMessage: err.Error()}
	}

	updatePlayerRequestObj := entities.PlayersUpdate{
		Status: constants.DB_STATUS_ACTIVE,
	}

	status := "failed"
	if code == activateCode {
		status = "Player activate successfully"
		err := service.redisDBAccess.DeleteRegistrationDataByKey(ctx, keyName)
		if err != nil {
			return responseData, &util.BadRequest{ErrMessage: err.Error()}
		}
		err = service.db.UpdatePlayerQuery(ctx, &updatePlayerRequestObj, playerCode)
		if err != nil {
			return responseData, err
		}
	}
	responseData.Data = nil
	responseData.Message = status
	responseData.RecordSet = nil
	return responseData, nil
}
