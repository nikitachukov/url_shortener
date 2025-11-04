package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/repository/dbrepo"
	"github.com/nikitachukov/url_shortener.git/internal/repository/memoryrepo"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func StartServer() {
	logger.InitLogger()

	configuration := config.NewParams()
	configuration.InitParams()

	if configuration.DSN != "" {
		service.InitRepo(dbrepo.NewDB(configuration.DSN))
	} else {
		service.InitRepo(memoryrepo.NewInMemory(configuration.FileStoragePath))
	}

	mux := chi.NewRouter()

	mux.Use(handler.LoggingHandlersMiddleware)
	mux.Use(handler.CustomDecompress)

	mux.Post("/", handler.MakeActionPost(configuration.BasePath))
	mux.Post("/api/shorten", handler.MakeActionPostAPI(configuration.BasePath))
	mux.Post("/api/shorten/batch", handler.MakeActionPostBatchAPI(configuration.BasePath))

	if configuration.BasePath != "" {
		mux.Get("/"+configuration.BasePath+"/{short}", handler.ActionGet)
	} else {
		mux.Get("/{short}", handler.ActionGet)
	}

	mux.Get("/ping", handler.Ping)

	serverPath := configuration.AppAddr

	server := &http.Server{
		Addr:    serverPath,
		Handler: mux,
	}

	logger.Log.Sugar().Infof("Starting server on: http://%s", serverPath)

	// Start server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {

			logger.Log.Sugar().Fatalf("HTTP server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop

	logger.Log.Sugar().Info("Shutdown signal received. Initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Sugar().Fatalf("HTTP server shutdown error: %v", err)
	}

	logger.Log.Sugar().Info("Server gracefully stopped.")
}
