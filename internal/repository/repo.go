package repository

import (
	"github.com/nikitachukov/url_shortener.git/internal/model"
)

const defaultCapacity = 1000

var (
	mapShortener = make(model.MapShortener, defaultCapacity)
)

func MapShorts() *model.MapShortener {
	return &mapShortener
}
