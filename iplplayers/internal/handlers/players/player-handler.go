package players

import (
	"golangliveprojects/iplplayers/pkg/util"
	"golangliveprojects/iplplayers/pkg/v1/services"

	"github.com/gin-gonic/gin"
)

type PlayerHandler interface {
	List(c *gin.Context)
	PlayerDetails(c *gin.Context)
	ListPlayerMatches(c *gin.Context)
	AddPlayer(c *gin.Context)
	UpdatePlayer(c *gin.Context)
}

type playerHandler struct {
	service services.PlayerService
}

func NewPlayerHandler(service services.PlayerService) PlayerHandler {
	return &playerHandler{service: service}
}

// method
func (srv *playerHandler) List(c *gin.Context) {
	resp, err := srv.service.List(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}

func (srv *playerHandler) PlayerDetails(c *gin.Context) {
	resp, err := srv.service.PlayerDetails(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}
func (srv *playerHandler) ListPlayerMatches(c *gin.Context) {
	resp, err := srv.service.ListPlayerMatches(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}

func (srv *playerHandler) AddPlayer(c *gin.Context) {
	resp, err := srv.service.AddPlayer(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}

func (srv *playerHandler) UpdatePlayer(c *gin.Context) {
	resp, err := srv.service.UpdatePlayer(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}
