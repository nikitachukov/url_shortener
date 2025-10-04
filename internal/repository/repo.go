package repository

import (
	"log"

	"github.com/nikitachukov/url_shortener.git/internal/model"
)

const defaultCapacity = 1000

type ShortenerRepo interface {
	FindShortURL(long string) (string, bool)
	GetLongURL(short string) (string, bool)
	Set(short, long string)
}

type MemoryRepo struct {
	m    model.MapShortener
	path string
}

func NewInMemory(path string) *MemoryRepo {
	repo := &MemoryRepo{m: make(model.MapShortener, defaultCapacity)}
	repo.Load(path)
	return repo
}

func (r *MemoryRepo) Load(path string) {
	r.path = path
	//	[
	//  {"uuid":"1","short_url":"4rSPg8ap","original_url":"http://yandex.ru"},
	//  {"uuid":"2","short_url":"edVPg3ks","original_url":"http://ya.ru"},
	//  {"uuid":"3","short_url":"dG56Hqxm","original_url":"http://practicum.yandex.ru"},
	//  ...
	//]
}

func (r *MemoryRepo) Save() {
	log.Println(r.path)
	//
}

func (r *MemoryRepo) FindShortURL(long string) (string, bool) {
	for short, l := range r.m {
		if l == long {
			return short, true
		}
	}
	return "", false
}

func (r *MemoryRepo) GetLongURL(short string) (string, bool) {
	l, ok := r.m[short]
	return l, ok
}

func (r *MemoryRepo) Set(short, long string) {
	r.m[short] = long
	r.Save()
}
