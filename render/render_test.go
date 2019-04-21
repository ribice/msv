package render_test

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ribice/msv/render"
)

type resp struct {
	FullName string `json:"full_name"`
	BornYear int    `json:"born_year"`
}

func TestJSON(t *testing.T) {

	cases := []struct {
		name        string
		wantStatus  int
		wantMessage string
		request     interface{}
	}{
		{
			name:        "Error decoding JSON",
			wantStatus:  500,
			wantMessage: "json: unsupported type: chan int",
			request:     make(chan int),
		},
		{
			name:        "Success",
			wantStatus:  200,
			request:     &resp{FullName: "Emir Ribic", BornYear: 1991},
			wantMessage: `{"full_name":"Emir Ribic","born_year":1991}`,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			render.JSON(w, tt.request)
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
