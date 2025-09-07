package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikitachukov/url_shortener.git/internal/handler"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
)

func TestHandlers(t *testing.T) {
	t.Run("test POST", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte("gopnik.win")))
		response := httptest.NewRecorder()
		handler.MainHandler(response, request)

		if response.Result().StatusCode != http.StatusCreated {
			t.Errorf("got %v, want %v", response.Result().Status, http.StatusCreated)
		}

	})

	t.Run("test GET", func(t *testing.T) {
		key := "newCode123"
		location := "http://example.com/page"
		(*repository.MapShorts())[key] = location

		request, _ := http.NewRequest("GET", "/newCode123", nil)
		response := httptest.NewRecorder()
		handler.MainHandler(response, request)

		result := response.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusTemporaryRedirect {
			t.Errorf("got %v, want %v", response.Result().Status, http.StatusTemporaryRedirect)
		}

		if result.Header.Get("Location") != location {
			t.Errorf("got %v, want %v", result.Header.Get("Location"), location)
		}

	})
}
