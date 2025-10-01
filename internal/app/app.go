package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func StartServer() {

	config.ParseParams()

	repo := repository.NewInMemory()
	service.InitRepo(repo)

	Mux := chi.NewRouter()
	Mux.Post("/", handler.ActionPost)

	if *config.BasePath != "" {
		Mux.Get("/"+*config.BasePath+"/{short}", handler.ActionGet)
	} else {
		Mux.Get("/{short}", handler.ActionGet)
	}

	serverPath := *config.AppAddr
	log.Printf("Starting server on: http://%s", serverPath)

	if err := http.ListenAndServe(serverPath, Mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stopped serving new connections.")
}
