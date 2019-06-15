package bauth_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ribice/msv/middleware/bauth"
)

func TestMiddleware(t *testing.T) {

	cases := []struct {
		name        string
		wantStatus  int
		wantMessage string
		user        string
		pass        string
	}{
		{
			name:        "Error decoding JSON",
			wantStatus:  401,
			wantMessage: "Unauthorized.",
		},
		{
			name:       "Success",
			wantStatus: 200,
			user:       "admin",
			pass:       "pass",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ba := bauth.New("admin", "pass", "DEFAULT")
			ts := httptest.NewServer(ba.MWFunc(testHandler()))
			defer ts.Close()
			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				t.Error(err)
			}
			req.SetBasicAuth(tt.user, tt.pass)
			cl := http.Client{}
			resp, err := cl.Do(req)
			if err != nil {
				t.Error(err)
			}
			defer resp.Body.Close()
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("Expected status: %v, got: %v", tt.wantStatus, resp.StatusCode)
			}
			bb := strings.TrimSpace(string(bodyBytes))
			if bb != tt.wantMessage {
				t.Errorf("Expected message: %v, got: %v", tt.wantMessage, bb)
			}

		})

	}
}

func testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
