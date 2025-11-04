package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
)

type Service struct {
	repo repository.ShortenerRepo
}

const asciiLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const lengthOfCode = 10

func NewService(r repository.ShortenerRepo) *Service {
	return &Service{repo: r}
}

func generateCode() (string, error) {
	b := make([]byte, lengthOfCode)
	for i := 0; i < lengthOfCode; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(asciiLetters))))
		if err != nil {
			return "", err
		}
		b[i] = asciiLetters[n.Int64()]
	}
	return string(b), nil
}

func (s *Service) ShortURL(long []byte) (string, bool, error) {
	longStr := strings.TrimSpace(string(long))
	if longStr == "" {
		return "", false, fmt.Errorf("empty URL is not allowed")
	}

	var short string

	code, err := generateCode()
	if err != nil {
		return "", false, fmt.Errorf("failed to generate code: %w", err)
	}

	if exists := s.repo.Set(code, longStr); !exists {
		short = code
		logger.Log.Sugar().Infof("Short url: %s set for long: %s", short, longStr)
		return short, false, nil
	} else {
		short, _ = s.repo.FindShortURL(longStr)
		logger.Log.Sugar().Infof("Short url: %s exsist! long: %s", short, longStr)
		return short, true, nil
	}
}

func (s *Service) GetLongURL(short string) (string, error) {
	if s == nil || s.repo == nil {
		return "", fmt.Errorf("repository is not initialized")
	}
	long, exists := s.repo.GetLongURL(short)
	if !exists {
		return "", fmt.Errorf("short url not found")
	}
	return long, nil
}

func (s *Service) Ping() bool {
	if s == nil || s.repo == nil {
		return false
	}
	return s.repo.Ping()
}
