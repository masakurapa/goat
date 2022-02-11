package goat

import "strings"

// Response is the type for manage response parameters.
type Response struct {
	// Status is the response status code
	Status int
	// Headers is the response headers
	Headers []H
	// Body is the response body
	Body string
}

func (r *Response) isJSON() bool {
	for _, header := range r.Headers {
		if header.Key == "Content-Type" && strings.HasPrefix(header.Value, contentTypeJSON) {
			return true
		}
	}
	return false
}
