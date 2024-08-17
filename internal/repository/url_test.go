package repository

import (
	"log"
	"os"
	"testing"
	"time"
	"url_shortener/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	InitDB()

	_, err := GetDB().Exec("CREATE TABLE IF NOT EXISTS url (origin varchar(255),create_at DATE, code varchar(255) PRIMARY KEY)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	code := m.Run()

	_, err = GetDB().Exec("DROP TABLE IF EXISTS url")
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	os.Exit(code)
}

func TestGetOriginByCode(t *testing.T) {
	testURL := model.URL{Origin: "https://example.com", Code: "abc123"}
	_, err := GetDB().Exec("INSERT INTO url (origin, code) VALUES (?, ?)", testURL.Origin, testURL.Code)
	if err != nil {
		t.Fatalf("Failed to insert test URL: %v", err)
	}

	result, err := GetOriginByCode(testURL.Code)

	assert.NoError(t, err)
	assert.Equal(t, testURL.Origin, result.Origin)
	assert.Equal(t, testURL.Code, result.Code)
}

func TestSave(t *testing.T) {
	testURL := model.URL{Origin: "https://example.com", Code: "xyz789", CreatedAt: time.Now()}

	code, err := Save(testURL)

	assert.NoError(t, err)
	assert.Equal(t, testURL.Code, code)

	row := GetDB().QueryRow("SELECT origin, code FROM url WHERE code = ?", testURL.Code)
	var origin, codeFromDB string
	err = row.Scan(&origin, &codeFromDB)
	if err != nil {
		t.Fatalf("Failed to retrieve URL from database: %v", err)
	}
	assert.Equal(t, testURL.Origin, origin)
	assert.Equal(t, testURL.Code, codeFromDB)
}
