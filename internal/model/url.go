package model

import "time"

type URL struct {
	Origin    string    `json:"origin"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}
