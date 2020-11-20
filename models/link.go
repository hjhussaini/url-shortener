package models

import (
	"time"

	"github.com/hjhussaini/url-shortener/database"
)

// swagger:model
type Link struct {
	ShortURL    string        `json:"-"`
	CustomURL   string        `json:"custom_url,omitempty"`
	OriginalURL string        `json:"original_url"`
	APIKey      string        `json:"api_key"`
	UserName    string        `json:"user_name,omitempty"`
	ExpireAt    time.Duration `json:"expire_at,omitempty"`
	TTL         int64         `json:"-"`
}

func (link *Link) Table() string {
	return "link"
}

func (link *Link) Fields() string {
	return "short_url, original_url, api_key, user_name, expire_at"
}

func (link *Link) Insert(session database.Session) error {
	return session.Insert(
		link.TTL,
		link.Table(),
		link.Fields(),
		link.ShortURL, link.OriginalURL, link.APIKey, link.UserName, link.ExpireAt,
	)
}
