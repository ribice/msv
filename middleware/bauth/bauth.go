package bauth

import (
	"crypto/subtle"
	"net/http"
)

// Service holds authorization middleware data
type Service struct {
	u     []byte
	p     []byte
	realm string
}

// New instantiates new auth middleware
func New(username, password, realm string) *Service {
	return &Service{u: []byte(username), p: []byte(password), realm: `Basic realm="` + realm + `"`}
}

// MWFunc implements http middleware for Basic Authentication
//  During initialization, username and password need to be provided which are compared afterwards in all requests.
func (s *Service) MWFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), s.u) != 1 || subtle.ConstantTimeCompare([]byte(pass), s.p) != 1 {
			w.Header().Set("WWW-Authenticate", s.realm)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
