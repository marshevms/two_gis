package main

import (
	"net/http"
	"time"

	"github.com/marshevms/two_gis/internal/logger"
)

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			uri := r.URL.String()
			method := r.Method
			logger.Info("start request: %s %s", method, uri)
			defer logger.Info("end request: %s %s (%s)", method, uri, time.Since(start))

			next.ServeHTTP(w, r)
		})
	}
}

