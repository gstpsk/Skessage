package util

import "net/http"

func StatusCodeResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
