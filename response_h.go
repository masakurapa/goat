package goat

func JsonResponse(status int, body string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json"})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func JsonResponseWithCharset(status int, body, charset string, headers ...H) Response {
	headers = append(headers, H{Key: "Content-Type", Value: "application/json; charset=" + charset})
	return Response{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}
