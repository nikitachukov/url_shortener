package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	//"github.com/nikitachukov/url_shortener.git/internal/model"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
)

type DBRepo struct {
	db  *sql.DB
	mu  sync.RWMutex
	dns string
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
	cnf, err := pgx.ParseConfig(dsn)
	if err != nil {
		fmt.Println("Error parsing DSN:", err)
	}

	mydsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Database)

	db, err := sql.Open("pgx", mydsn)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New("file://migrations", mydsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Log.Sugar().Info("No changes found in migrations")
		}

	}

	repo := &DBRepo{
		db:  db,
		dns: mydsn,
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
	if err := r.db.QueryRow("select key from shortener where original_url=$1 limit 1", long).Scan(&short); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		return "", false
	}
	return short, true
}

func (r *DBRepo) GetLongURL(short string) (string, bool) {
	var long string
	if err := r.db.QueryRow("select original_url from shortener where key=$1 limit 1", short).Scan(&long); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		return "", false
	}
	return long, true

}

func (r *DBRepo) Set(short, long string) {
	r.db.Exec("INSERT INTO shortener (key, original_url) VALUES ($1, $2)", short, long)
}
