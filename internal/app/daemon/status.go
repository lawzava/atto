package daemon

import (
	"atto/pkg/format"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

const (
	statusOk = "ok"
)

func statusHandler(c *gin.Context) {
	response := statusResponse{
		Timestamp: format.Time.Build(time.Now()),
		Status:    statusOk,
	}

	c.JSON(http.StatusOK, response)
}
