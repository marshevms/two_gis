package middleware

import (
	"net/http"
	"runtime"
	"time"

	"github.com/marshevms/two_gis/internal/logger"
)

func checkMethod(method string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				logger.Infof("method not allowed: %s", r.Method)
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func Get() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return checkMethod(http.MethodGet)(next)
	}
}

func Post() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return checkMethod(http.MethodPost)(next)
	}
}

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			uri := r.URL.String()
			method := r.Method
			logger.Infof("start request: %s %s", method, uri)
			defer func() {
				logger.Infof("end request: %s %s (%s)", method, uri, time.Since(start))
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Recover() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					logger.Errorf("panic: %v\n%s", err, buf)

					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
