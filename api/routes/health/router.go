package health

import (
	"github.com/gorilla/mux"
)

// Map mapeia as rotas à partir do path raiz
func Map(router *mux.Router) {
	router.HandleFunc("/health", Handle)
}
