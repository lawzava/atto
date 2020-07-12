package daemon

import (
	"atto/pkg/handling"
	"net/http"
	"time"
)

type statusResponse struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

const (
	statusOk = "ok"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Timestamp: handling.Time.Build(time.Now()),
		Status:    statusOk,
	}

	err := handling.RenderJSON(w, http.StatusOK, response)
	if err != nil {
		// TODO: Handle errors appropriately via logger
		panic(err)
	}
}
