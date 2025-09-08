package router

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
)

func StartServer() {

	key := "newCode123"
	location := "http://example.com/page"
	(*repository.MapShorts())[key] = location

	mux := chi.NewRouter()
	mux.Post("/", handler.ActionPost)
	mux.Get("/{short}", handler.ActionGet)

	serverPath := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)
	log.Printf("Starting server on: http://%s", serverPath)

	if err := http.ListenAndServe(serverPath, mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stopped serving new connections.")
}
