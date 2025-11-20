package middleware

import "net/http"

type WriteWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WriteWrapper) WriteHeader(StatusCode int) {
	w.ResponseWriter.WriteHeader(StatusCode)
	w.StatusCode = StatusCode
}
