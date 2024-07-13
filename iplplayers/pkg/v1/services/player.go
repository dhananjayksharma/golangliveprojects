package services

import (
	"golangliveprojects/iplplayers/pkg/v1/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerService interface {
	List(c *gin.Context) (responses.Response, error)
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
	return responseData, nil
}
