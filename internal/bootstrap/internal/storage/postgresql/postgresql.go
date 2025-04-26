package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"url-shortener/internal/bootstrap/internal/config"
)

type Storage struct {
	Db *sql.DB
}

func New(config *config.Storage) (*Storage, error) {
	const op = "storage.postgres.New"

	port, err := strconv.Atoi(config.Port)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.Host, port, config.User, config.Pass, config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//stmt, err := db.Prepare(`
	//CREATE TABLE IF NOT EXISTS urls (
	//   id INTEGER PRIMARY KEY AUTO_INCREMENT,
	//	alias TEXT UNIQUE NOT NULL,
	//	url TEXT UNIQUE NOT NULL);
	//Create INDEX IF NOT EXISTS idx_alias ON urls (alias);
	//`)
	//if err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	//}
	//_, err = stmt.Exec()
	//if err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	//}

	return &Storage{Db: db}, nil
}
