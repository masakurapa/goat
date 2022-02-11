package goat

import "net/http"

func JsonResponse(status int, body string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: contentTypeJSON + "; charset=utf-8"})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func JsonResponseWithoutCharset(status int, body, charset string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: contentTypeJSON})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func NoContent(headers ...H) Response {
	return Response{
		Status:  http.StatusNoContent,
		Body:    "",
		Headers: headers,
	}
}
