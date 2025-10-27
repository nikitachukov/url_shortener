package repository

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

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
	r.path = filename
	fileData, err := os.ReadFile(r.path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, &r.m)
	if err != nil {
		return err
	}
	r.currentID = len(r.m)

	return nil

}

func (r *MemoryRepo) Save() {
	if r.path != "" {

		fileData, err := json.Marshal(r.m)
		if err != nil {
			log.Println(err)
		}
		err = os.WriteFile(r.path, fileData, 0644)
		if err != nil {
			log.Println(err)
		}
		return

	}

}

func (r *MemoryRepo) FindShortURL(long string) (string, bool) {
	for _, l := range r.m {
		if l.OriginalURL == long {

			return l.ShortURL, true
		}
	}
	return "", false
}

func (r *MemoryRepo) GetLongURL(short string) (string, bool) {
	item, ok := r.m[short]
	if ok {
		return item.OriginalURL, true
	} else {
		return "", false
	}
}

func (r *MemoryRepo) Set(short, long string) {
	var item model.Item
	r.currentID++
	item.UUID = strconv.Itoa(r.currentID)
	item.ShortURL = short
	item.OriginalURL = long
	r.m[short] = item
	r.Save()
}
