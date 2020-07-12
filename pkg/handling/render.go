package handling

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, httpStatus int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encode JSON response: %w", err)
	}

	return nil
}
