package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

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
	var item model.Item

	r.mu.Lock()
	r.path = filename

	file, err := os.Open(r.path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r.currentID = 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line %d: %s\n", r.currentID, line)
		err := json.Unmarshal([]byte(line), &item)
		if err != nil {
			log.Fatal(err)
		}
		r.m[item.ShortURL] = item
		r.currentID++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.mu.Unlock()
	return nil

}

func (r *MemoryRepo) Save(short string) {
	path := r.path
	item := r.m[short]

	if path != "" {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()

		line, _ := json.Marshal(item)

		if _, err := file.WriteString("\n" + string(line)); err != nil {
			fmt.Println("Ошибка записи:", err)
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
	r.mu.Lock()
	item.UUID = strconv.Itoa(r.currentID)
	item.ShortURL = short
	item.OriginalURL = long
	r.m[short] = item
	r.currentID++
	r.Save(short)
	r.mu.Unlock()
}
