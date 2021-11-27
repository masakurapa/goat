package goat

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"

	"github.com/masakurapa/qstringer"
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
	// Header is the request header
	Header H
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

	if r.Header != nil {
		for k, v := range r.Header {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func (r *Request) url(serv *httptest.Server) (string, error) {
	url := serv.URL + r.Path
	if r.Query == nil {
		return url, nil
	}

	query, err := qstringer.Encode(r.Query)
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

func (q *Q) join() (string, error) {
	if len(*q) == 0 {
		return "", nil
	}

	params := url.Values{}
	for k, v := range *q {
		r := reflect.ValueOf(v)
		switch r.Kind() {
		case reflect.Bool:
			params.Add(k, fmt.Sprintf("%v", r.Bool()))
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			params.Add(k, fmt.Sprintf("%d", r.Int()))
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			params.Add(k, fmt.Sprintf("%d", r.Uint()))
		case reflect.Float32, reflect.Float64:
			params.Add(k, fmt.Sprintf("%d", r.Float()))
		case reflect.Slice:
			sk := k
			if !strings.HasSuffix(k, "[]") {
				sk += "[]"
			}

			for i := 0; i < r.Len(); i++ {
				sv := r.Index(0)
				switch sv.Kind() {

				case reflect.Bool:
					params.Add(k, fmt.Sprintf("%v", r.Bool()))
				case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
					params.Add(k, fmt.Sprintf("%d", r.Int()))
				case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
					params.Add(k, fmt.Sprintf("%d", r.Uint()))
				case reflect.Float32, reflect.Float64:
					params.Add(k, fmt.Sprintf("%d", r.Float()))
				case reflect.String:
					params.Add(k, r.String())
				}
			}

			params.Add(sk, fmt.Sprintf("%d", r.Float()))
		case reflect.String:
			params.Add(k, r.String())
		// 	float32, float64:
		// 	params.Add(k, fmt.Sprintf("%d", vt))
		default:
			// TODO: メッセージ
			return "", fmt.Errorf("invalid query string type %q", r.Kind().String())
		}
	}

	return "?" + params.Encode(), nil
}
