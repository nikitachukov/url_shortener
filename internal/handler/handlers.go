package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nikitachukov/url_shortener.git/internal/repository"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func ActionGet(res http.ResponseWriter, req *http.Request) {
	shortParam := req.URL.Path[1:]
	longURL, ok := (*repository.MapShorts())[shortParam]
	if !ok {
		log.Printf("Unable to find longURL URL for short: %s: status: %d", shortParam, http.StatusBadRequest)
		http.Error(res, "Unable to find longURL URL for short", http.StatusBadRequest)
	}

	log.Println("Map of short links: ", repository.MapShorts())
	res.Header().Add("Location", longURL)

	res.WriteHeader(http.StatusTemporaryRedirect)
	log.Println("Full header: ", res.Header())

}

func ActionPost(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Unable to read body: status: %d", http.StatusBadRequest)
		http.Error(res, "Unable to read body", http.StatusBadRequest)
	}
	defer req.Body.Close()

	shortURL, err := service.ShortURL(body)
	if err != nil {
		log.Printf("Unable to shorten URL: status: %d", http.StatusBadRequest)
		http.Error(res, "Unable to shorten URL", http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)

	_, err = res.Write([]byte(fmt.Sprintf("http://%s/%s", req.Host, shortURL)))
	if err != nil {
		log.Printf("Unexpected exception: status: %d", http.StatusInternalServerError)
		http.Error(res, "Unexpected exception: ", http.StatusInternalServerError)
	}

}
