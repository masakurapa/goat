package goat

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	emptyFunc = func(*testing.T) {}
)

// TestCase is the type for manage API test cases
type TestCase struct {
	// Name is the name of the test case
	Name string
	// SetUp is the function that runs before the API is executed
	SetUp func(*testing.T)
	// TearDown is the function that runs after the API is executed
	TearDown func(*testing.T)
	// Request is the request parameter of the API
	Request Request
	// Response is the response parameter of the API
	Response Response
	// AfterFuncs is the slice to set up a function to perform arbitrary validation after API execution
	AfterFuncs []func(*testing.T)
}

// H is the type of the request and response headers
type H struct {
	Key   string
	Value string
}

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
	if f == nil {
		return
	}
	r.setUpBeforeTest = f
}

// TearDownAfterTest set up a function that all test cases will run after execution
func (r *T) TearDownAfterTest(f func(*testing.T)) {
	if f == nil {
		return
	}
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

			for _, f := range tc.AfterFuncs {
				f(t)
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

	endpoint := fmt.Sprintf("%s %s", request.Method, request.Path)

	if actual.StatusCode != expected.Status {
		t.Errorf("%q is not returns code returens %d, want %d", endpoint, actual.StatusCode, expected.Status)
	}

	for _, h := range expected.Headers {
		if _, ok := actual.Header[h.Key]; !ok {
			// TODO: メッセージ
			t.Errorf("%q is not returns header %q", endpoint, h.Key)
			continue
		}
		if a := actual.Header.Get(h.Key); a != h.Value {
			// TODO: メッセージ
			t.Errorf("%q header %q returens %q, want %q", endpoint, h.Key, a, h.Value)
		}
	}

	responseBody, err := io.ReadAll(actual.Body)
	if err != nil {
		t.Fatal(responseBody)
	}
	// TODO: ResponseBodyに余計なスペースとか含まれる可能性を考慮してcompact化、trimをしてから検証する
	if actual := string(responseBody); actual != string(expected.Body) {
		t.Errorf("%q body returens %s, want %s", endpoint, actual, string(expected.Body))
	}
}
