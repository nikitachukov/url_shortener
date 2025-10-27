package dbRepo

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	//"github.com/nikitachukov/url_shortener.git/internal/model"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
)

type DbRepo struct {
	db *sql.DB
	mu sync.RWMutex
}

func (r *DbRepo) Ping() bool {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return false
	} else {
		return true
	}

}

func NewDB(dsn string) *DbRepo {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	repo := &DbRepo{
		db: db,
	}

	logger.Log.Sugar().Infof("DB connection established")

	return repo
}

func (r *DbRepo) Load(filename string) error {
	return nil
}

func (r *DbRepo) Save() {

}

func (r *DbRepo) FindShortURL(long string) (string, bool) {
	return "", false
}

func (r *DbRepo) GetLongURL(short string) (string, bool) {
	return "", false
}

func (r *DbRepo) Set(short, long string) {

}
