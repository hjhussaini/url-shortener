package common

import (
	"encoding/json"
	"io"
)

// ReadJSON deserializes the object from JSON string
func ReadJSON(reader io.Reader, object interface{}) error {
	return json.NewDecoder(reader).Decode(object)
}

// WriteJSON
