package players

import (
	"golangliveprojects/iplplayers/pkg/util"
	"golangliveprojects/iplplayers/pkg/v1/services"

	"github.com/gin-gonic/gin"
)

type PlayerHandler interface {
	List(c *gin.Context)
	PlayerDetails(c *gin.Context)
	ListMatches(c *gin.Context)
}

type playerHandler struct {
	service services.PlayerService
}

func NewPlayerHandler(service services.PlayerService) PlayerHandler {
	return &playerHandler{service: service}
}

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
func (srv *playerHandler) ListMatches(c *gin.Context) {
	resp, err := srv.service.ListMatches(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp, resp.Message)
}
