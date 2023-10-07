package data

import "time"

type Watch struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Brand     string    `json:"brand,omitempty"`
	Model     string    `json:"model,omitempty"`
	DialColor string    `json:"dial_color"`
	StrapType string    `json:"strap_type"`
	Price     float64   `json:"price"`
}
