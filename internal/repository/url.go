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

func CheckOriginExists(origin string) bool {
	var exists bool
	row := GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM url WHERE origin = ?)", origin)
	err := row.Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func GetCodeByOrigin(origin string) (string, error) {
	var url model.URL
	row := GetDB().QueryRow("SELECT code FROM url WHERE origin = ?", origin)
	if err := row.Scan(&url.Code); err != nil {
		if err == sql.ErrNoRows {
			return url.Code, fmt.Errorf("Url with origin %s is not found: %w", origin, err)
		}
		log.Printf("Error rettrieving URL by origin %s: %v", origin, err)
		return url.Code, err
	}
	return url.Code, nil
}
