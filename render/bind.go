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
func Bind(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return Respond(w, InternalF("error decoding json: %v", err))
	}

	if binder, ok := v.(Binder); ok {
		err := binder.Bind()
		if err != nil {
			return Respond(w, BadRequestF("error binding json: %v", err))
		}
	}
	return nil
}
