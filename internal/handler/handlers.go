package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func ActionGet(res http.ResponseWriter, req *http.Request) {

	shortParam := chi.URLParam(req, "short")
	if shortParam == "" {
		shortParam = req.URL.Path[1:]
	}

	longURL, err := service.GetLongURL(shortParam)
	if err != nil {
		logger.Log.Sugar().Infof("Unable to find longURL URL for short: %s: status: %d", shortParam, http.StatusBadRequest)
		http.Error(res, "Unable to find longURL URL for short", http.StatusBadRequest)
		return
	}

	res.Header().Add("Location", longURL)

	res.WriteHeader(http.StatusTemporaryRedirect)

}

func MakeActionPost(basePath string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Log.Sugar().Errorf("Unable to read body: status: %d", http.StatusBadRequest)
			http.Error(res, "Unable to read body", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		shortURL, err := service.ShortURL(body)
		if err != nil {
			logger.Log.Sugar().Errorf("Unable to shorten URL: status: %d", http.StatusBadRequest)
			http.Error(res, "Unable to shorten URL", http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusCreated)
		if basePath == "" {
			_, err = res.Write([]byte(fmt.Sprintf("http://%s/%s", req.Host, shortURL)))
		} else {
			_, err = res.Write([]byte(fmt.Sprintf("http://%s/%s/%s", req.Host, basePath, shortURL)))
		}
		if err != nil {
			logger.Log.Sugar().Errorf("Unexpected exception: status: %d", http.StatusInternalServerError)
			http.Error(res, "Unexpected exception: ", http.StatusInternalServerError)
			return
		}
	}
}
