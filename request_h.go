package goat

import (
	"net/http"
)

// ConnectRequest is the helper that returns the configuration for sending a CONNECT request
func ConnectRequest(path string, query Q, header H) Request {
	return Request{
		Method: http.MethodConnect,
		Path:   path,
		Query:  query,
		Header: header,
	}
}

// DeleteRequest is the helper that returns the configuration for sending a DELETE request
func DeleteRequest(path, body string, query Q, header H) Request {
	return Request{
		Method: http.MethodDelete,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// DeleteJsonRequest is the helper that returns the configuration for sending a DELETE request with a JSON format body.
func DeleteJsonRequest(path, body string, query Q, header H) Request {
	return Request{
		Method: http.MethodDelete,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// GetRequest is the helper that returns the configuration for sending a GET request
func GetRequest(path string, query Q, header H) Request {
	return Request{
		Method: http.MethodGet,
		Path:   path,
		Query:  query,
		Header: header,
	}
}

// HeadRequest is the helper that returns the configuration for sending a HEAD request
func HeadRequest(path string, query Q, header H) Request {
	return Request{
		Method: http.MethodGet,
		Path:   path,
		Query:  query,
		Header: header,
	}
}

// OptionsRequest is the helper that returns the configuration for sending a OPTIONS request
func OptionsRequest(path string, query Q, header H) Request {
	return Request{
		Method: http.MethodOptions,
		Path:   path,
		Query:  query,
		Header: header,
	}
}

// PatchRequest is the helper that returns the configuration for sending a PATCH request
func PatchRequest(path, body string, query Q, header H) Request {
	return Request{
		Method: http.MethodPatch,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// PatchJsonRequest is the helper that returns the configuration for sending a PATCH request with a JSON format body.
func PatchJsonRequest(path, body string, query Q, header H) Request {
	return Request{
		Method: http.MethodPatch,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request
func PostRequest(path, body string, query Q, header H) Request {
	return Request{
		Method: http.MethodPost,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// PostRequest is the helper that returns the configuration for sending a POST request with a JSON format body.
func PostJsonRequest(path, body string, query Q, header H) Request {
	header["Content-Type"] = "application/json"
	return Request{
		Method: http.MethodPost,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// PutRequest is the helper that returns the configuration for sending a PUT request
func PutRequest(path, body string, query Q, header H) Request {
	header["Content-Type"] = "application/json"
	return Request{
		Method: http.MethodPut,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// PutJsonRequest is the helper that returns the configuration for sending a PUT request with a JSON format body.
func PutJsonRequest(path, body string, query Q, header H) Request {
	header["Content-Type"] = "application/json"
	return Request{
		Method: http.MethodPut,
		Path:   path,
		Query:  query,
		Body:   body,
		Header: header,
	}
}

// TraceRequest is the helper that returns the configuration for sending a TRACE request
func TraceRequest(path string, query Q, header H) Request {
	return Request{
		Method: http.MethodTrace,
		Path:   path,
		Query:  query,
		Header: header,
	}
}
