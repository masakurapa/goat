package goat

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/masakurapa/qstring"
)

// Request is the type for manage request parameters.
type Request struct {
	// Method is the request method
	//
	// see https://github.com/golang/go/blob/master/src/net/http/method.go for available methods.
	Method string
	// Path is the request path.
	//
	// since we will be using an internally started test server, the scheme, domain, and port are not necessary.
	Path string
	// Query is the query string parameter
	Query Q
	// Headers is the request headers
	Headers []H
	// Body is the request body
	Body string
}

func (r *Request) makeRequest(serv *httptest.Server) (*http.Request, error) {
	url, err := r.url(serv)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.Method, url, r.body())
	if err != nil {
		return nil, err
	}

	for _, h := range r.Headers {
		req.Header.Add(h.Key, h.Value)
	}

	return req, nil
}

func (r *Request) url(serv *httptest.Server) (string, error) {
	url := serv.URL + r.Path
	if r.Query == nil {
		return url, nil
	}

	query, err := qstring.Encode(r.Query)
	if err != nil {
		return "", err
	}
	return url + query, nil
}

func (r *Request) body() io.Reader {
	if !r.hasBody() {
		return nil
	}
	return bytes.NewReader([]byte(r.Body))
}

func (r *Request) hasBody() bool {
	return r.Method == http.MethodDelete ||
		r.Method == http.MethodPatch ||
		r.Method == http.MethodPost ||
		r.Method == http.MethodPut
}

// Q is the type of the query string parameters
type Q map[string]interface{}

// ArrayQ is a type of query string in array format
type ArrayQ []interface{}
