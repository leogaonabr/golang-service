package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/leogaonabr/golang-service/api/routes"
	"github.com/leogaonabr/golang-service/config"
	"github.com/leogaonabr/golang-service/util"
)

// StartServer starts the REST API server
func StartServer() *http.Server {
	logger := util.GetLogger()
	logger.Info("starting golang-template-app HTTP server")

	router := mux.NewRouter()
	logger.Info("starting mapping routes")
	routes.Map(router)

	logger.Info("routes mapped, now setting up middlewares")
	router.Use(handlers.RecoveryHandler())
	router.Use(util.Tracer)
	router.Use(util.Metrificator)

	addr := fmt.Sprintf(":%d", config.GetPort())
	logger.Infof("server configured to listen at addr %s", addr)

	// creates the server instance
	srv := http.Server{
		Handler:      router,
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// starts the http listener on another goroutine. The daemon routine will be kept blocked with a os.Signal channel
	go func(srv *http.Server) {
		err := srv.ListenAndServe()
		if err != nil {
			logger.Errorf("error listening on port: %s", err.Error())
			os.Exit(1)
		}
	}(&srv)

	logger.Info("golang-template-app HTTP server up and running")
	return &srv
}
