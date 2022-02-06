package goat

// Response is the type for manage response parameters.
type Response struct {
	// Status is the response status code
	Status int
	// Headers is the response headers
	Headers []H
	// Body is the response body
	Body string
}
