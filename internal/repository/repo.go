package repository

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/nikitachukov/url_shortener.git/internal/logger"
	"github.com/nikitachukov/url_shortener.git/internal/model"
)

type ShortenerRepo interface {
	FindShortURL(long string) (string, bool)
	GetLongURL(short string) (string, bool)
	Set(short, long string)
}

type MemoryRepo struct {
	m         model.MapShortener
	path      string
	currentID int
	mu        sync.RWMutex
}

func NewInMemory(filename string) *MemoryRepo {

	repo := &MemoryRepo{m: make(model.MapShortener, 0)}
	repo.currentID = 0

	if filename != "" {
		err := repo.Load(filename)
		if err != nil {
			return repo
		}
	}
	return repo
}

func (r *MemoryRepo) Load(filename string) error {
	var items []model.Item

	r.mu.Lock()
	r.path = filename

	data, err := os.ReadFile(r.path)
	if err != nil {
		logger.Log.Sugar().Infof("Error reading file: %s", err)
	}

	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		logger.Log.Sugar().Infof("Error parcing file: %s", err)
	}

	for _, item := range items {
		r.m[item.ShortURL] = item
		r.currentID++
	}

	logger.Log.Sugar().Infof("Loaded %d items from %s", len(items), filename)

	r.mu.Unlock()

	return nil

}

func (r *MemoryRepo) Save() {
	path := r.path

	if path != "" {

		var result strings.Builder
		result.WriteString("[\n")
		var i int
		for _, item := range r.m {
			i++
			jsonBytes, _ := json.Marshal(item)
			result.WriteString("  ")
			result.Write(jsonBytes)
			if i < r.currentID {
				result.WriteString(",")
			}
			result.WriteString("\n")
		}

		result.WriteString("]")

		err := os.WriteFile(r.path, []byte(result.String()), 0644)
		if err != nil {
			return
		}

	}

}

func (r *MemoryRepo) FindShortURL(long string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, l := range r.m {
		if l.OriginalURL == long {
			return l.ShortURL, true
		}
	}
	return "", false
}

func (r *MemoryRepo) GetLongURL(short string) (string, bool) {
	r.mu.RLock()
	item, ok := r.m[short]
	r.mu.RUnlock()
	if ok {
		return item.OriginalURL, true
	} else {
		return "", false
	}
}

func (r *MemoryRepo) Set(short, long string) {
	var item model.Item
	r.currentID++
	r.mu.Lock()
	item.UUID = strconv.Itoa(r.currentID)
	item.ShortURL = short
	item.OriginalURL = long
	r.m[short] = item
	r.Save()
	r.mu.Unlock()
}
