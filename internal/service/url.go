package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"
	"url_shortener/internal/dto"
	"url_shortener/internal/model"
	"url_shortener/internal/repository"
)

func GetOriginByCode(code string) (model.URL, error) {
	if code == "" {
		return model.URL{}, errors.New("code cannot be empty")
	}
	return repository.GetOriginByCode(code)
}

func Save(urlDTO dto.UrlDTO) (string, error) {
	if err := validateUrl(urlDTO.Origin); err != nil {
		return "", err
	}

	code, err := generateUniqueCode(15)
	if err != nil {
		return "", err
	}

	url := model.URL{Origin: urlDTO.Origin, Code: code}
	return repository.Save(url)
}

func validateUrl(rawUrl string) error {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return errors.New("invalid URL format")
	}
	return nil
}

func generateUniqueCode(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(bytes)
	if len(encoded) < length {
		return encoded, nil
	}
	return encoded[:length], nil
}
