package example_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/masakurapa/goat"
)

func TestExample(t *testing.T) {
	// set `*gin.Engine` to `goat.New()`
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
