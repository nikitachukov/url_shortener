package repository

import (
	"github.com/nikitachukov/url_shortener.git/internal/model"
)

const defaultCapacity = 1000

type ShortenerRepo interface {
	FindShortURL(long string) (string, bool)
	GetLongURL(short string) (string, bool)
	Set(short, long string)
}

type MemoryRepo struct {
	m model.MapShortener
}

func NewInMemory() *MemoryRepo {
	return &MemoryRepo{m: make(model.MapShortener, defaultCapacity)}
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
}
