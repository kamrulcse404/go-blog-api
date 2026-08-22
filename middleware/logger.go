package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

func (sw *statusWriter) Write(b []byte) (int, error) {
	if sw.statusCode == 0 {
		sw.statusCode = http.StatusOK
	}

	return sw.ResponseWriter.Write(b)
}

func Logger(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		sw := &statusWriter{
			ResponseWriter: w,
			statusCode:     0,
		}

		next.ServeHTTP(sw, r)
		duration := time.Since(start)
		fmt.Println(r.Method, r.URL.Path, sw.statusCode, duration)
	}

	return http.HandlerFunc(logger)
}
