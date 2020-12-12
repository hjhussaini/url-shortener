package api

import "net/http"

// swagger:route GET /{short_link}
func (links *Links) RedirectURL(writer http.ResponseWriter, request *http.Request) {
}
