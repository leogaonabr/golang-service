package health

import (
	"encoding/json"
	"net/http"
)

// Handle a rota GET /health
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	response, _ := json.Marshal(map[string]interface{}{
		"status": "up",
	})

	w.Write(response)
	return
}
