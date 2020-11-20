package api

import "github.com/hjhussaini/url-shortener/database"

type Link struct {
	Session database.Session
}

func NewLinkAPIs(session database.Session) *Link {
	return &Link{Session: session}
}
