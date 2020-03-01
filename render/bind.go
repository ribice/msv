package render

import (
	"encoding/json"
	"net/http"
)

// Binder binds json
type Binder interface {
	Bind() error
}

// Bind binds JSON request into interface v, and validates the request
//  If an error occurs during json unmarshalling, http 500 is responded to client
//  If an error occurs during binding json, http 400 is responded to client
func Bind(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	if binder, ok := v.(Binder); ok {
		return binder.Bind()
	}
	return nil
}
