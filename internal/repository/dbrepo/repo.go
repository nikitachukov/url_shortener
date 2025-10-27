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
	var short string
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	row := r.db.QueryRowContext(ctx, "select key from shortener where original_url=$1 limit 1", long)
	err := row.Scan(&short)
	if err != nil {
		panic(err)
	}
	if short != "" {
		return short, true
	} else {
		return "", false
	}
}

func (r *DBRepo) GetLongURL(short string) (string, bool) {
	var long string
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	row := r.db.QueryRowContext(ctx, "select original_url from shortener where key=$1 limit 1", short)
	err := row.Scan(&long)
	if err != nil {
		panic(err)
	}
	if short != "" {
		return long, true
	} else {
		return "", false
	}
}

func (r *DBRepo) Set(short, long string) {
	_, err := r.db.Exec("INSERT INTO shortener (key, original_url) VALUES ($1, $2)", short, long)
	if err != nil {
	}
}
