package model

import "time"

type URL struct {
	Id        int64     `json:"id"`
	Origin    string    `json:"origin"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}
