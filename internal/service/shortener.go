package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/repository"
)

const asciiLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const lengthOfCode = 10

var Repo repository.ShortenerRepo

func InitRepo(r repository.ShortenerRepo) {
	Repo = r
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

func ShortURL(long []byte) (string, error) {
	longStr := strings.TrimSpace(string(long))
	if longStr == "" {
		return "", fmt.Errorf("empty URL is not allowed")
	}

	if shortURL, ok := Repo.FindShortURL(longStr); ok {
		logger.Log.Sugar().Infof("Short url: %s got for long: %s", shortURL, longStr)
		return shortURL, nil
	}

	var short string

	for attempts := 0; attempts < 255; attempts++ {
		code, err := generateCode()
		if err != nil {
			return "", fmt.Errorf("failed to generate code: %w", err)
		}
		if _, exists := Repo.GetLongURL(code); !exists {
			short = code
			Repo.Set(short, longStr)
			break
		}
	}

	logger.Log.Sugar().Infof("Short url: %s set for long: %s", short, longStr)
	return short, nil
}

func GetLongURL(short string) (string, error) {
	if Repo == nil {
		return "", fmt.Errorf("repository is not initialized")
	}
	long, exists := Repo.GetLongURL(short)
	if !exists {
		return "", fmt.Errorf("short url not found")
	}
	return long, nil
}

func Ping() bool {
	return Repo.Ping()
}
