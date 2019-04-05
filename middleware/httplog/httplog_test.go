package httplog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	})
	myHandlerWithError = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusUnavailableForLegalReasons), http.StatusUnavailableForLegalReasons)
	})
)

func TestDefault(t *testing.T) {
	svc := New("prefix")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/should/be/stdout/", nil)
	req.Header.Set("X-Real-IP", "100.100.100.100")
	svc.MWFunc(myHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusOK)
	expect(t, res.Body.String(), "test")
}

func TestIgnoredURIs(t *testing.T) {
	svc := New("prefix", "/ignore")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ignore", nil)
	req.RequestURI = "/ignore"
	svc.MWFunc(myHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusOK)
	expect(t, res.Body.String(), "test")
}

func TestXForwardedFor(t *testing.T) {
	svc := New("prefix")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/should/be/stdout/", nil)
	req.Header.Add("X-Forwarded-For", "100.100.100.100")
	svc.MWFunc(myHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusOK)
	expect(t, res.Body.String(), "test")
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
