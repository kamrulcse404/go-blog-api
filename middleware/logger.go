package middleware

import (
	"fmt"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.ResponseWriter.WriteHeader(code)
	sw.statusCode = code
}

func Logger(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter, r *http.Request) {

		sw := &statusWriter{
			ResponseWriter: w,
			statusCode: http.StatusOK,
		}

		next.ServeHTTP(sw, r)
		fmt.Println(r.Method, r.URL.Path, sw.statusCode)
	}

	return http.HandlerFunc(logger)
}
