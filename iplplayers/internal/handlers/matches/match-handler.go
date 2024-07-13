package matches

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Match list"})
}
