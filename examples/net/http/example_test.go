package example_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
	// set `http.Handler` to `goat.New()`
	goat.New(handler()).Run(t, []goat.TestCase{
		{
			Name:     "example test case",
			SetUp:    func(*testing.T) {},
			TearDown: func(*testing.T) {},
			Request:  goat.GetJsonRequest("/example"),
			Response: goat.JsonResponse(http.StatusOK, `{"ping":"pong"}`),
		},
	})
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/example", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, `{"ping":"pong"}`)
	}))
	return mux
}
