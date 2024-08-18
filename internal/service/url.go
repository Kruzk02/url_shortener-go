package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/url"
	"time"
	"url_shortener/internal/dto"
	"url_shortener/internal/model"
	"url_shortener/internal/repository"

	"github.com/redis/go-redis/v9"
)

func GetOriginByCode(code string) (model.URL, error) {
	if code == "" {
		return model.URL{}, errors.New("code cannot be empty")
	}

	cached, err := repository.GetRedis().Get(context.Background(), code).Result()
	if err == nil {
		return model.URL{
			Origin: cached,
			Code:   code,
		}, nil
	} else if err != redis.Nil {
		return model.URL{}, err
	}

	url, err := repository.GetOriginByCode(code)
	if err != nil {
		return model.URL{}, err
	}

	err = repository.GetRedis().Set(context.Background(), code, url.Origin, time.Hour*2).Err()
	if err != nil {
		return url, err
	}

	return url, nil
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
	if err := repository.GetRedis().Set(context.Background(), url.Code, url.Origin, time.Hour*2).Err(); err != nil {
		log.Printf("Error caching URL in Redis: %v", err)
	}

	if _, err := repository.Save(url); err != nil {
		return "", err
	}

	return code, nil
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
