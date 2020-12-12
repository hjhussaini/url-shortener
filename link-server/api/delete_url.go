package api

import "net/http"

// swagger:route DELETE /links/{short_link}
func (links *Links) DeleteURL(writer http.ResponseWriter, request *http.Request) {
}
