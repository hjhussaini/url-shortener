package models

import "time"

// swagger:model
type Link struct {
	ShortURL    string        `json:"-"`
	CustomURL   string        `json:"custom_url,omitempty"`
	OriginalURL string        `json:"original_url"`
	APIKey      string        `json:"api_key"`
	UserName    string        `json:"user_name,omitempty"`
	ExpireAt    time.Duration `json:"expire_at,omitempty"`
}
