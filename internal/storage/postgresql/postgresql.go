package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"url-shortener/internal/config"
	"url-shortener/internal/storage"
)

type Storage struct {
	Db *sql.DB
}

func FactoryStorage(config config.Storage) (storage.Storage, error) {
	const op = "storage.postgres.New"

	port, err := strconv.Atoi(config.GetPort())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.GetHost(), port, config.GetUser(), config.GetPassword(), config.GetDBName())
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{Db: db}, nil
}

func (s Storage) SaveUrl(url, alice string) error {
	return errors.New(url + alice)
}
