package recovery_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ribice/msv/middleware/recovery"
)

var (
	myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bar"))
	})
	myPanicHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("this did not work")
	})
)

func TestNoConfigGood(t *testing.T) {
	r := recovery.New("prefix")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/should/not/panic/", nil)
	r.MWFunc(myHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusOK)
	expect(t, res.Body.String(), "bar")
}

func TestDefaultPanic(t *testing.T) {
	r := recovery.New("prefix")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/should/panic/", nil)
	r.MWFunc(myPanicHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusInternalServerError)
	expect(t, strings.TrimSpace(res.Body.String()), strings.TrimSpace(http.StatusText(http.StatusInternalServerError)))
}

func TestCustomPanicHandler(t *testing.T) {
	r := recovery.New("prefix")

	customHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	})
	r.SetPanicHandler(customHandler)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/should/400/", nil)
	r.MWFunc(myPanicHandler).ServeHTTP(res, req)

	expect(t, res.Code, http.StatusBadRequest)
	expect(t, res.Body.String(), "Bad Request")
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
