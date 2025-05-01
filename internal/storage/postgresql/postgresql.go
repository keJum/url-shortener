package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strconv"
	"url-shortener/internal/config"
	storageErr "url-shortener/internal/lib/storage/errors"
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
	const op = "storage.postgres.SaveUrl"

	stmt, err := s.Db.Prepare("INSERT INTO urls(\"url\", \"alias\") VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(url, alice)
	if err != nil {

		var postgresErr *pq.Error
		if errors.As(err, &postgresErr) && postgresErr.Code.Name() == "unique_violation" {
			return fmt.Errorf("%s: %w", op, storageErr.ErrUrlExists)
		}
	}
	return err
}

func (s Storage) GetUrl(alice string) (string, error) {
	const op = "storage.postgres.GetUrl"
	var url string

	stmt, err := s.Db.Prepare("SELECT url FROM urls WHERE alias = $1")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRow(alice).Scan(&url)
	if err != nil {
		var postgresErr *pq.Error
		if errors.As(err, &postgresErr) && postgresErr.Code.Name() == "not_found" {
			return "", fmt.Errorf("%s: %w", op, storageErr.ErrUrlNotFound)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, storageErr.ErrUrlNotFound)
		} else {
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}
	return url, nil
}
