package repository

import (
	"database/sql"
	"fmt"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (u URLRepository) IsUrlExist(url string) bool {
	err := u.db.QueryRow("select count(*) from urls wher original_url = $1", url)
	if err == nil {
		return true
	}
	return false
}

func (u URLRepository) AddUrl(originalUrl string, shortUrl string) error {
	_, err := u.db.Exec("INSERT INTO urls (original_url, short_code) VALUES ($1, $2)", originalUrl, shortUrl)

	if err != nil {
		return err
	}
	return nil
}

func (u *URLRepository) GetNextId() (int64, error) {
	var id int64
	err := u.db.QueryRow("SELECT nextval('urls_id_seq')").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get next id: %w", err)
	}
	return id, nil
}
