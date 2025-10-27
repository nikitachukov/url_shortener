package dbrepo

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	//"github.com/nikitachukov/url_shortener.git/internal/model"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
)

type DBRepo struct {
	db *sql.DB
	mu sync.RWMutex
}

func (r *DBRepo) Ping() bool {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return false
	} else {
		return true
	}

}

func NewDB(dsn string) *DBRepo {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	repo := &DBRepo{
		db: db,
	}

	logger.Log.Sugar().Infof("DB connection established")

	return repo
}

func (r *DBRepo) Load() {
}

func (r *DBRepo) Save() {

}

func (r *DBRepo) FindShortURL(long string) (string, bool) {
	return "", false
}

func (r *DBRepo) GetLongURL(short string) (string, bool) {
	return "", false
}

func (r *DBRepo) Set(short, long string) {

}
