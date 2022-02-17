package goat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	contentTypeJson = "application/json"
)

var (
	emptyFunc = func(*testing.T) {}
)

// TestCase is the type for manage API test cases
type TestCase struct {
	// Name is the name of the test case
	Name string
	// SetUp is the function that runs before the API is executed.
	//
	// This function performs preprocessing such as migration, test data creation, etc.
	SetUp func(*testing.T)
	// TearDown is the function that runs after the API is executed
	//
	// This function perform post-processing such as deleting test data, deleting tables, etc.
	TearDown func(*testing.T)
	// CustomTestFuncs is a set of functions used when you want to perform validation other
	// than HTTP status, response header, and response body validation.
	//
	// Each `CustomTestFunc` works in series.
	//
	// If you use `t.Fatal()`, `t.Fail()`, etc. in a `CustomTestFunc`,
	// the process will be interrupted and the subsequent `CustomTestFunc` will not be processed.
	CustomTestFuncs []CustomTestFunc
	// Request is the request parameter of the API
	Request Request
	// Response is the response parameter of the API
	Response Response
}

// H is the type of the request and response headers
type H struct {
	Key   string
	Value string
}

// CustomTestFunc is a type of function that implements its own validation
type CustomTestFunc func(t *testing.T, res *http.Response)

// New returns a client for API testing
func New(handler http.Handler) *T {
	return &T{
		handler:           handler,
		setUpBeforeTest:   emptyFunc,
		tearDownAfterTest: emptyFunc,
	}
}

// T is the type for manage API tests.
type T struct {
	handler           http.Handler
	setUpBeforeTest   func(*testing.T)
	tearDownAfterTest func(*testing.T)
}

// SetUpBeforeTest set up a function that all test cases will run before execution
func (r *T) SetUpBeforeTest(f func(*testing.T)) {
	r.setUpBeforeTest = f
}

// TearDownAfterTest set up a function that all test cases will run after execution
func (r *T) TearDownAfterTest(f func(*testing.T)) {
	r.tearDownAfterTest = f
}

// Run executes the test case
func (r *T) Run(t *testing.T, testCases []TestCase) {
	if r.setUpBeforeTest != nil {
		r.setUpBeforeTest(t)
	}
	if r.tearDownAfterTest != nil {
		defer r.tearDownAfterTest(t)
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.SetUp != nil {
				tc.SetUp(t)
			}
			if tc.TearDown != nil {
				defer tc.TearDown(t)
			}

			resp := r.send(t, tc.Request)
			r.assertResponse(t, tc.Request, resp, tc.Response)

			for _, f := range tc.CustomTestFuncs {
				t.Run("CustomTestFunc", func(t *testing.T) {
					f(t, resp)
				})
			}
		})
	}
}

func (r *T) send(t *testing.T, request Request) *http.Response {
	serv := httptest.NewServer(r.handler)
	defer serv.Close()

	req, err := request.makeRequest(serv)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func (r *T) assertResponse(t *testing.T, request Request, actual *http.Response, expected Response) {
	defer actual.Body.Close()

	endpoint := request.Method + " " + request.Path

	if actual.StatusCode != expected.Status {
		t.Errorf("[%v] status code returns %d, want %d", endpoint, actual.StatusCode, expected.Status)
	}

	for _, h := range expected.Headers {
		if _, ok := actual.Header[h.Key]; !ok {
			t.Errorf("[%v] key %q does not exist in header", endpoint, h.Key)
			continue
		}

		val := actual.Header.Get(h.Key)
		a := strings.ToLower(val)
		e := strings.ToLower(h.Value)
		if a != e {
			t.Errorf("[%v] %q is set to the key %q in the header, want %q", endpoint, h.Key, val, h.Value)
		}
	}

	responseBody, err := io.ReadAll(actual.Body)
	if err != nil {
		t.Fatal(responseBody)
	}

	// if the status is 204, not validate response body
	if expected.Status == http.StatusNoContent {
		return
	}

	a := strings.TrimSpace(string(responseBody))
	e := strings.TrimSpace(expected.Body)

	// if the expected value is JSON, json.Compact and compare it
	if expected.isJson() {
		isErr := false
		if s, err := r.compactJson(a); err != nil {
			t.Errorf("response body is not JSON format. error: %s", err)
			isErr = true
		} else {
			a = s
		}

		if s, err := r.compactJson(e); err != nil {
			t.Errorf("expected response body is not JSON format. error: %s", err)
			isErr = true
		} else {
			e = s
		}

		if isErr {
			return
		}
	}

	if a != e {
		t.Errorf("[%v] body returens %s, want %s", endpoint, a, e)
	}
}

func (r *T) compactJson(s string) (string, error) {
	dist := &bytes.Buffer{}
	if err := json.Compact(dist, []byte(s)); err != nil {
		return "", err
	}
	return dist.String(), nil
}
