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

func JsonResponse(status int, body string, headers ...H) Response {
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func JsonResponseWithCharset(status int, body, charset string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}
