package goat

import (
	"net/http"
)

// ConnectRequest is the helper that returns the configuration for sending a CONNECT request
func ConnectRequest(path string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodConnect,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}

// DeleteRequest is the helper that returns the configuration for sending a DELETE request
func DeleteRequest(path, body string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodDelete,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// DeleteJsonRequest is the helper that returns the configuration for sending a DELETE request with a JSON format body.
func DeleteJsonRequest(path, body string, query Q, headers ...H) Request {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	return Request{
		Method:  http.MethodDelete,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// GetRequest is the helper that returns the configuration for sending a GET request
func GetRequest(path string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodGet,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}

// HeadRequest is the helper that returns the configuration for sending a HEAD request
func HeadRequest(path string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodGet,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}

// OptionsRequest is the helper that returns the configuration for sending a OPTIONS request
func OptionsRequest(path string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodOptions,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}

// PatchRequest is the helper that returns the configuration for sending a PATCH request
func PatchRequest(path, body string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodPatch,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// PatchJsonRequest is the helper that returns the configuration for sending a PATCH request with a JSON format body.
func PatchJsonRequest(path, body string, query Q, headers ...H) Request {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	return Request{
		Method:  http.MethodPatch,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request
func PostRequest(path, body string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodPost,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request with a JSON format body.
func PostJsonRequest(path, body string, query Q, headers ...H) Request {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	return Request{
		Method:  http.MethodPost,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// PutRequest is the helper that returns the configuration for sending a PUT request
func PutRequest(path, body string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodPut,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// PutJsonRequest is the helper that returns the configuration for sending a PUT request with a JSON format body.
func PutJsonRequest(path, body string, query Q, headers ...H) Request {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=utf8"})
	return Request{
		Method:  http.MethodPut,
		Path:    path,
		Query:   query,
		Body:    body,
		Headers: headers,
	}
}

// TraceRequest is the helper that returns the configuration for sending a TRACE request
func TraceRequest(path string, query Q, headers ...H) Request {
	return Request{
		Method:  http.MethodTrace,
		Path:    path,
		Query:   query,
		Headers: headers,
	}
}
