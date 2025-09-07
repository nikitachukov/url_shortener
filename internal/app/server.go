package router

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MainHandler)

	serverPath := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)
	log.Printf("Starting server on: http://%s", serverPath)

	if err := http.ListenAndServe(serverPath, mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stopped serving new connections.")
}
