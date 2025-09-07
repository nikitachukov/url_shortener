package service

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/nikitachukov/url_shortener.git/internal/repository"
)

const asciiLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const lengthOfCode = 10

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

	m := repository.MapShorts()

	for shortURL, longURL := range *m {
		if longStr == longURL {
			log.Printf("Short url: %s got for long: %s", shortURL, longStr)
			return shortURL, nil
		}
	}

	var short string
	for {
		code, err := generateCode()
		if err != nil {
			return "", fmt.Errorf("failed to generate code: %w", err)
		}
		if _, exists := (*m)[code]; !exists {
			short = code
			(*m)[short] = longStr
			break
		}
	}

	log.Printf("Short url: %s set for long: %s", short, longStr)
	return short, nil
}
