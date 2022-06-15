package example_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
	tt := goat.New(handler())

	tt.SetUpBeforeTest(func(t *testing.T) { fmt.Println("run SetUpBeforeTest") })
	tt.TearDownAfterTest(func(t *testing.T) { fmt.Println("run TearDownAfterTest") })

	tt.Run(t, []goat.TestCase{
		{
			Name:     "test1",
			SetUp:    func(*testing.T) { fmt.Println("run SetUp test1") },
			TearDown: func(*testing.T) { fmt.Println("run TearDown test1") },
			Request:  goat.GetJsonRequest("/example"),
			Response: goat.JsonResponse(http.StatusOK, `{"ping":"pong"}`),
			CustomTestFuncs: []goat.CustomTestFunc{
				{Name: "CustomTestFunc1", Func: func(t *testing.T, res *http.Response) { fmt.Println("run CustomTestFunc1") }},
				{Name: "CustomTestFunc2", Func: func(t *testing.T, res *http.Response) { fmt.Println("run CustomTestFunc2") }},
			},
		},
		{
			Name:     "test2",
			SetUp:    func(*testing.T) { fmt.Println("run SetUp test2") },
			TearDown: func(*testing.T) { fmt.Println("run TearDown test2") },
			Request:  goat.GetJsonRequest("/example"),
			Response: goat.JsonResponse(http.StatusOK, `{"ping":"pong"}`),
			CustomTestFuncs: []goat.CustomTestFunc{
				{Func: func(t *testing.T, res *http.Response) { fmt.Println("run CustomTestFunc3") }},
				{Func: func(t *testing.T, res *http.Response) { fmt.Println("run CustomTestFunc4") }},
			},
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
