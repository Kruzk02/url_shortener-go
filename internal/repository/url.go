package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"url_shortener/internal/model"
)

func GetOriginByCode(code string) (model.URL, error) {
	var url model.URL
	row := GetDB().QueryRow("SELECT origin, code FROM url WHERE code = ?", code)
	if err := row.Scan(&url.Origin, &url.Code); err != nil {
		if err == sql.ErrNoRows {
			return url, fmt.Errorf("URL with code %s not found: %w", code, err)
		}
		log.Printf("Error retrieving URL by code %s: %v", code, err)
		return url, err
	}
	return url, nil
}

func Save(url model.URL) (string, error) {
	tx, err := GetDB().Begin()
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO url (origin,code,create_at) VALUES (?,?,?)", url.Origin, url.Code, time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to save URL: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit tranction: %w", err)
	}

	return url.Code, nil
}
