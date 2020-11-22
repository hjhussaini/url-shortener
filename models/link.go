package models

import (
	"time"

	"github.com/hjhussaini/url-shortener/database"
)

// swagger:model
type Link struct {
	UserID    string        `json:"user_id"`
	UserName  string        `json:"user_name,omitempty"`
	ShortURL  string        `json:"-"`
	CustomURL string        `json:"custom_url,omitempty"`
	LongURL   string        `json:"long_url"`
	ExpireAt  time.Duration `json:"expire_at,omitempty"`
	TTL       int64         `json:"-"`
}

func (link *Link) Table() string {
	return "link"
}

func (link *Link) Fields() string {
	return "user_id, short_url, long_url, expire_at"
}

func (link *Link) Insert(session database.Session) error {
	return session.Insert(
		link.TTL,
		link.Table(),
		link.Fields(),
		link.UserID, link.ShortURL, link.LongURL, link.ExpireAt,
	)
}
