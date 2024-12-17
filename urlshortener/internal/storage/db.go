package storage

import (
	"context"
	"database/sql"
	"fmt"
	"urlShort/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	pool *sql.DB
}

func New(connString string) (*DB, error) {
	pool, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

func AddUrl(ctx context.Context, db *DB, urls models.Urls) error {
	tx, err := db.pool.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO url(long_url, short_url) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, urls.Url, urls.ShortUrl)
	if err != nil {
		return err
	}
	fmt.Println(res.RowsAffected())

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetUrls(ctx context.Context, db *DB, url models.Url) ([]models.Urls, error) {
	rows, err := db.pool.QueryContext(ctx, `
		SELECT long_url
		FROM url
		WHERE short_url = ?;
	`, url.Url)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []models.Urls
	for rows.Next() {
		var u models.Urls
		err := rows.Scan(
			&u.Url,
			&u.ShortUrl,
		)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return urls, nil
}
