package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikitachukov/url_shortener.git/internal/config"
	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
	"github.com/nikitachukov/url_shortener.git/internal/repository/dbrepo"
	"github.com/nikitachukov/url_shortener.git/internal/repository/memoryrepo"
	"github.com/nikitachukov/url_shortener.git/internal/service"
)

func TestHandlers(t *testing.T) {
	var repo repository.ShortenerRepo
	configuration := config.NewParams()
	configuration.InitParams()

	if configuration.DSN != "" {
		repo = dbrepo.NewDB(configuration.DSN)
		service.InitRepo(repo)
	} else {
		repo = memoryrepo.NewInMemory(configuration.FileStoragePath)
		service.InitRepo(repo)
	}
	t.Run("test POST", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte("gopnik.win")))
		response := httptest.NewRecorder()
		handler.MakeActionPost(configuration.BasePath)(response, request)
		result := response.Result()
		defer result.Body.Close()
		if result.StatusCode != http.StatusCreated {
			t.Errorf("got %v, want %v", result.Status, http.StatusCreated)
		}

	})

	t.Run("test GET", func(t *testing.T) {
		key := "newCode123"
		location := "http://example.com/page"
		repo.Set(key, location)

		request, _ := http.NewRequest("GET", "/"+key, nil)
		response := httptest.NewRecorder()
		handler.ActionGet(response, request)

		result := response.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusTemporaryRedirect {
			t.Errorf("got %v, want %v", result.Status, http.StatusTemporaryRedirect)
		}

		if result.Header.Get("Location") != location {
			t.Errorf("got %v, want %v", result.Header.Get("Location"), location)
		}

	})
}
