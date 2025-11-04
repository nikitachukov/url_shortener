package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	_render "github.com/go-chi/render"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

type APIShortenReq struct {
	URL string `json:"url"`
}
type APIShortenRes struct {
	Result string `json:"result"`
}

type APIShortenBatchReq struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type APIShortenBatchRes struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func MakeActionGet(svc *service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		shortParam := chi.URLParam(req, "short")
		if shortParam == "" {
			shortParam = req.URL.Path[1:]
		}

		longURL, err := svc.GetLongURL(shortParam)
		if err != nil {
			logger.Log.Sugar().Infof("Unable to find longURL URL for short: %s: status: %d", shortParam, http.StatusBadRequest)
			http.Error(res, "Unable to find longURL URL for short", http.StatusBadRequest)
			return
		}

		res.Header().Add("Location", longURL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func MakeActionPost(svc *service.Service, basePath string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Log.Sugar().Errorf("Unable to read body: status: %d", http.StatusBadRequest)
			http.Error(res, "Unable to read body", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		shortURL, exsist, err := svc.ShortURL(body)
		if err != nil {
			logger.Log.Sugar().Errorf("Unable to shorten URL: status: %d", http.StatusBadRequest)
			http.Error(res, "Unable to shorten URL", http.StatusBadRequest)
			return
		}
		if !exsist {
			res.WriteHeader(http.StatusCreated)
		} else {
			res.WriteHeader(http.StatusConflict)
		}

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

func MakeActionPostAPI(svc *service.Service, basePath string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var data APIShortenReq
		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		defer req.Body.Close()

		shortURL, exsist, err := svc.ShortURL([]byte(data.URL))
		if err != nil {
			logger.Log.Sugar().Errorf("Unable to shorten URL: status: %d", http.StatusBadRequest)
			http.Error(res, "Unable to shorten URL", http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json")

		if !exsist {
			res.WriteHeader(http.StatusCreated)
		} else {
			res.WriteHeader(http.StatusConflict)
		}

		if basePath == "" {

			_render.JSON(res, req, APIShortenRes{Result: fmt.Sprintf("http://%s/%s", req.Host, shortURL)})
		} else {
			_render.JSON(res, req, APIShortenRes{Result: fmt.Sprintf("http://%s/%s/%s", req.Host, basePath, shortURL)})
		}

	}
}

func PingHandler(svc *service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if svc.Ping() {
			res.WriteHeader(http.StatusOK)
			return
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func MakeActionPostBatchAPI(svc *service.Service, basePath string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var (
			reqData []APIShortenBatchReq
			resData []APIShortenBatchRes
		)

		res.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(req.Body).Decode(&reqData)

		if err != nil {
			logger.Log.Sugar().Errorf("Unable to shorten URL: status: %d", http.StatusBadRequest)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		defer req.Body.Close()

		for _, item := range reqData {
			shortURL, _, err := svc.ShortURL([]byte(item.OriginalURL))

			if err != nil {
				logger.Log.Sugar().Errorf("Unable to shorten URL: status: %d", http.StatusInternalServerError)
				res.WriteHeader(http.StatusInternalServerError)
				return
			}

			if basePath == "" {
				resData = append(resData, APIShortenBatchRes{CorrelationID: item.CorrelationID, ShortURL: fmt.Sprintf("http://%s/%s", req.Host, shortURL)})
			} else {
				resData = append(resData, APIShortenBatchRes{CorrelationID: item.CorrelationID, ShortURL: fmt.Sprintf("http://%s/%s/%s", req.Host, basePath, shortURL)})
			}
		}

		res.WriteHeader(http.StatusCreated)

		_render.JSON(res, req, resData)

	}
}
