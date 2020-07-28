package routes

import (
	"github.com/leogaonabr/golang-service/api/routes/health"
	"github.com/gorilla/mux"
)

// Map mapeia as rotas à partir do path raiz
func Map(router *mux.Router) {
	health.Map(router)
}
