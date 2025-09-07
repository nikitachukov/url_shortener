package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/nikitachukov/url_shortener.git/internal/repository"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func mainHandlerActionGet(res http.ResponseWriter, req *http.Request) {
	longUrl, ok := (*repository.MapShorts())[req.RequestURI[1:]]
	if !ok {
		log.Printf("Unable to find longUrl URL for short: %s: status: %d", req.RequestURI[1:], http.StatusBadRequest)
		http.Error(res, "Unable to find longUrl URL for short", http.StatusBadRequest)
	}

	log.Println("Map of short links: ", repository.MapShorts())
	res.Header().Add("Location", longUrl)

	res.WriteHeader(http.StatusTemporaryRedirect)
	log.Println("Full header: ", res.Header())

}

func mainHandlerActionPost(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Unable to read body: status: %d", http.StatusBadRequest)
		http.Error(res, "Unable to read body", http.StatusBadRequest)
	}
	defer req.Body.Close()

	shortUrl, err := service.ShortURL(body)
	if err != nil {
		log.Printf("Unable to shorten URL: status: %d", http.StatusBadRequest)
		http.Error(res, "Unable to shorten URL", http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)

	_, err = res.Write([]byte("http://" + req.Host + "/" + shortUrl))
	if err != nil {
		log.Printf("Unexpected exception: status: %d", http.StatusInternalServerError)
		http.Error(res, "Unexpected exception: ", http.StatusInternalServerError)
	}

}

func MainHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodPost:
		mainHandlerActionPost(res, req)

	case http.MethodGet:
		mainHandlerActionGet(res, req)

	default:
		log.Printf("Method not implemented: %s, status: %d", req.Method, http.StatusNotImplemented)
		http.Error(res, "BadRequest", http.StatusBadRequest)

	}

}
