package app

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func StartServer() {
	logger.InitLogger()

	configuration := config.NewParams()
	configuration.InitParams()

	repo := repository.NewInMemory()
	service.InitRepo(repo)

	Mux := chi.NewRouter()

	Mux.Use(handler.HandlersLog)

	Mux.Post("/", handler.MakeActionPost(*configuration.BasePath))
	Mux.Post("/api/shorten", handler.MakeActionPostApi(*configuration.BasePath))

	if *configuration.BasePath != "" {
		Mux.Get("/"+*configuration.BasePath+"/{short}", handler.ActionGet)
	} else {
		Mux.Get("/{short}", handler.ActionGet)
	}

	serverPath := *configuration.AppAddr
	logger.Log.Sugar().Infof("Starting server on: http://%s", serverPath)

	if err := http.ListenAndServe(serverPath, Mux); !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Sugar().Fatalf("HTTP server error: %v", err)
	}

	logger.Log.Sugar().Info("Stopped serving new connections.")
}
