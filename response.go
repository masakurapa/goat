package goat

// Response is the type for manage response parameters.
type Response struct {
	// Status is the response status code
	Status int
	// Header is the response headers
	Header H
	// Body is the response body
	Body string
}

func JsonResponse(status int, body string, header H) Response {
	return JsonResponseWithCharset(status, body, "utf8", header)
}

func JsonResponseWithCharset(status int, body, charset string, header H) Response {
	header["Content-Type"] = "application/json; charset=" + charset
	return Response{
		Status: status,
		Body:   body,
		Header: header,
	}
}
