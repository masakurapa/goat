# goat
`goat` is golang API testing library.

# Usage

## using `gin`

If you are using [gin-gonic/gin](https://github.com/gin-gonic/gin), set `*gin.Engine` to `goat.New()`.

```go
package example_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
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

func handler() *gin.Engine {
	gin.DefaultWriter = io.Discard
	r := gin.Default()
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	return r
}
```

## using `echo`

If you are using [labstack/echo](https://github.com/labstack/echo), set `*echo.Echo` to `goat.New()`.

```go
package example_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
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

func handler() *echo.Echo {
	e := echo.New()
	e.GET("/example", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"ping": "pong"})
	})
	return e
}
```

## using `net/http` package

If you are using [net/http](https://pkg.go.dev/net/http), set `http.Handler` to `goat.New()`.

```go
package example_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
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
```

# SetUp and TearDown

Use `SetUp` and `TearDown` if you want the process to run before or after the test is executed.

Use `SetUpBeforeTest` and `TearDownAfterTest` if you want to run the process before and after all test cases are executed.

## Example

```go
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
			Request:  goat.GetJsonRequest("/ok"),
			Response: goat.JsonResponse(http.StatusOK, `{"status":"ok"}`),
		},
		{
			Name:     "test2",
			SetUp:    func(*testing.T) { fmt.Println("run SetUp test2") },
			TearDown: func(*testing.T) { fmt.Println("run TearDown test2") },
			Request:  goat.GetJsonRequest("/ok"),
			Response: goat.JsonResponse(http.StatusOK, `{"status":"ok"}`),
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
```

The result of running this test will look like this.

```sh
$ go test example_test.go -v
=== RUN   TestExample
run SetUpBeforeTest
=== RUN   TestExample/test1
run SetUp test1
run TearDown test1
=== RUN   TestExample/test2
run SetUp test2
run TearDown test2
run TearDownAfterTest
--- PASS: TestExample (0.00s)
    --- PASS: TestExample/test1 (0.00s)
    --- PASS: TestExample/test2 (0.00s)
PASS
ok      command-line-arguments  0.251s
```
