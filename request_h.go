package goat

import (
	"net/http"
)

// ConnectRequest is the helper that returns the configuration for sending a CONNECT request
func ConnectRequest(path string, headers ...H) Request {
	return Request{
		Method:  http.MethodConnect,
		Path:    path,
		Headers: headers,
	}
}

// DeleteRequest is the helper that returns the configuration for sending a DELETE request
func DeleteRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodDelete,
		Path:    path,
		Body:    body,
		Headers: headers,
	}
}

// DeleteJsonRequest is the helper that returns the configuration for sending a DELETE request with a JSON format body.
func DeleteJsonRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodDelete,
		Path:    path,
		Body:    body,
		Headers: appendJsonHeader(headers),
	}
}

// GetRequest is the helper that returns the configuration for sending a GET request
func GetRequest(path string, headers ...H) Request {
	return Request{
		Method:  http.MethodGet,
		Path:    path,
		Headers: headers,
	}
}

// HeadRequest is the helper that returns the configuration for sending a HEAD request
func HeadRequest(path string, headers ...H) Request {
	return Request{
		Method:  http.MethodGet,
		Path:    path,
		Headers: headers,
	}
}

// OptionsRequest is the helper that returns the configuration for sending a OPTIONS request
func OptionsRequest(path string, headers ...H) Request {
	return Request{
		Method:  http.MethodOptions,
		Path:    path,
		Headers: headers,
	}
}

// PatchRequest is the helper that returns the configuration for sending a PATCH request
func PatchRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPatch,
		Path:    path,
		Body:    body,
		Headers: headers,
	}
}

// PatchJsonRequest is the helper that returns the configuration for sending a PATCH request with a JSON format body.
func PatchJsonRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPatch,
		Path:    path,
		Body:    body,
		Headers: appendJsonHeader(headers),
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request
func PostRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPost,
		Path:    path,
		Body:    body,
		Headers: headers,
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request with a JSON format body.
func PostJsonRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPost,
		Path:    path,
		Body:    body,
		Headers: appendJsonHeader(headers),
	}
}

// PutRequest is the helper that returns the configuration for sending a PUT request
func PutRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPut,
		Path:    path,
		Body:    body,
		Headers: headers,
	}
}

// PutJsonRequest is the helper that returns the configuration for sending a PUT request with a JSON format body.
func PutJsonRequest(path, body string, headers ...H) Request {
	return Request{
		Method:  http.MethodPut,
		Path:    path,
		Body:    body,
		Headers: appendJsonHeader(headers),
	}
}

// TraceRequest is the helper that returns the configuration for sending a TRACE request
func TraceRequest(path string, headers ...H) Request {
	return Request{
		Method:  http.MethodTrace,
		Path:    path,
		Headers: headers,
	}
}

func appendJsonHeader(headers []H) []H {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	headers = append(headers, H{Key: "Accept", Value: "application/json"})
	return headers
}
