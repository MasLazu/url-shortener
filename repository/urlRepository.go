package repository

import (
	"context"
	"errors"
	"url-shortener/config"
	"url-shortener/model"
)

type UrlRepository struct {
	DB *config.Database
}

func NewUrlRepository(DB *config.Database) *UrlRepository {
	return &UrlRepository{
		DB: DB,
	}
}

func (r *UrlRepository) Save(ctx context.Context, url model.Url) error {
	sql := "INSERT INTO url (id, url) VALUES ($1, $2)"
	if _, err := r.DB.Conn.ExecContext(ctx, sql, url.ID, url.Url); err != nil {
		return err
	}

	return nil
}

func (r *UrlRepository) FindById(ctx context.Context, id string) (model.Url, error) {
	var url model.Url
	sql := "SELECT id, url FROM url WHERE id = $1"
	row, err := r.DB.Conn.QueryContext(ctx, sql, id)
	if err != nil {
		return url, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&url.ID, &url.Url); err != nil {
			return url, err
		}
		return url, nil
	}
	return url, errors.New("not found")
}
