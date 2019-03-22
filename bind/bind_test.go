package bind_test

import (
	"errors"
	"net/http/httptest"
	"testing"
)

func (r *req) Bind() error {
	if r.Age < 18 {
		return errors.New("age must be greater than 18")
	}
	r.Name = "Mike Johansson"
	return nil
}

type req struct {
	Name string
	Age  int
}

func TestJSON(t *testing.T) {

	cases := []struct {
		name        string
		wantStatus  int
		wantMessage string
		request     string
		wantStr     *req
	}{
		{
			name:        "Error decoding JSON",
			wantStatus:  500,
			wantMessage: "error decoding json",
			request:     "",
		},
		{
			name:        "Error binding data",
			wantStatus:  400,
			wantMessage: "error binding data:",
			request:     "",
		},
		{
			name:        "Success",
			wantStatus:  400,
			wantMessage: "error binding data:",
			request:     "",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer()
		})

	}
}
