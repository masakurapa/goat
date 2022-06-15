package example_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
	// set `*echo.Echo` to `goat.New()`
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
