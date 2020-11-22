package app

import (
	"net/http"
	"mime"
)

func ValidateJsonType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "" {
			NewErrorResponse("content-Type header must be application/json", http.StatusBadRequest).Write(w)
			return
		}
		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			NewErrorResponse("malformed Content-Type header", http.StatusBadRequest).Write(w)
			return
		}

		if mt != "application/json" {
			NewErrorResponse("content-Type header must be application/json", http.StatusUnsupportedMediaType).Write(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
