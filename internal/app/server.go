package router

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
)

func StartServer() {

	config.ParseParams()

	//key := "newCode123"
	//location := "http://example.com/page"
	//(*repository.MapShorts())[key] = location

	Mux := chi.NewRouter()
	Mux.Post("/", handler.ActionPost)
	Mux.Get("/"+*config.BasePath+"/{short}", handler.ActionGet)

	serverPath := fmt.Sprintf("%s", *config.AppAddr)
	log.Printf("Starting server on: http://%s", serverPath)

	if err := http.ListenAndServe(serverPath, Mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stopped serving new connections.")
}
