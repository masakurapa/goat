package goat

import "net/http"

// JsonResponse returns the structure of the expected value of the JSON response.
// charset is specified as utf-8.
func JsonResponse(status int, body string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: contentTypeJson + "; charset=utf-8"})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

// JsonResponse returns the structure of the expected value of the JSON response.
// No charset is specified.
func JsonResponseWithoutCharset(status int, body, charset string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: contentTypeJson})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

// NoContent returns the structure of the expected value of 204(No Content)
func NoContent(headers ...H) Response {
	return Response{
		Status:  http.StatusNoContent,
		Body:    "",
		Headers: headers,
	}
}
