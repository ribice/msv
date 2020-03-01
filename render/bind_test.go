package render_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ribice/msv/render"
)

func (r *req) Bind() error {
	if r.Age < 18 {
		return errors.New("age must be greater than 18")
	}
	r.Name = "Mike Johansson"
	return nil
}

type req struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestBind(t *testing.T) {
	cases := map[string]struct {
		wantStatus  int
		wantMessage string
		request     string
		wantStr     req
	}{
		"Error decoding JSON": {
			wantStatus:  500,
			wantMessage: `{"message":"error decoding json: invalid character '}' after object key"}`,
			request:     `{"invalid:json"}`,
		},
		"Error binding data": {
			wantStatus:  400,
			wantMessage: `{"message":"error binding json: age must be greater than 18"}`,
			request:     `{"name":"Emir", "age":15}`,
		},
		"Success": {
			wantStatus: 200,
			request:    `{"name":"Emir", "age":28}`,
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/", bytes.NewBufferString(tt.request))
			if err != nil {
				t.Error(err)
			}
			var th *testhandler
			th.ServeHTTP(w, req)
			defer w.Result().Body.Close()
			bodyBytes, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				t.Error(err)
			}
			if w.Code != tt.wantStatus {
				t.Errorf("Expected status: %v, got: %v", tt.wantStatus, w.Code)
			}
			bb := strings.TrimSpace(string(bodyBytes))
			if bb != tt.wantMessage {
				t.Errorf("Expected message: %v, got: %v", tt.wantMessage, bb)
			}
		})
	}
}

type testhandler struct{}

func (*testhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rq req
	if err := render.Bind(w, r, &rq); err != nil {
		return
	}
	w.WriteHeader(200)
}
