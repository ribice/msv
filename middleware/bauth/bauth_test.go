package bauth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestMiddleware(t *testing.T) {

	cases := []struct {
		name        string
		wantStatus  int
		wantMessage string
		request     string
		user        string
		pass        string
	}{
		{
			name:        "Error decoding JSON",
			wantStatus:  500,
			wantMessage: "error decoding json: invalid character '}' after object key",
			request:     `{"invalid:json"}`,
		},
		{
			name:        "Error binding data",
			wantStatus:  400,
			wantMessage: "error binding data: age must be greater than 18",
			request:     `{"name":"Emir", "age":15}`,
		},
		{
			name:       "Success",
			wantStatus: 200,
			request:    `{"name":"Emir", "age":28}`,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			j := jwt.New("testingsecret", 10, "HS256", tt.sess)
			ts := httptest.NewServer(j.MWFunc(testHandler()))
			defer ts.Close()
		})

	}
}

type testhandler struct{}

func (*testhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
}
