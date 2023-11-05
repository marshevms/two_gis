package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux  *http.ServeMux
	ip   string
	port int64

	routeExist  bool
	middlewares []func(next http.Handler) http.Handler
}

func New(ip string, port int64) *Server {
	return &Server{
		ip:   ip,
		port: port,
		mux:  http.NewServeMux(),
	}
}

func (s Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.routeExist = true
	s.mux.Handle(pattern, chain(http.HandlerFunc(handler), s.middlewares...))
}

// SetMiddleware устанавливает миддлвари, должно быть вызвано до установки роута
func (s Server) AddMiddleware(middlewares ...func(next http.Handler) http.Handler) {
	if s.routeExist {
		panic("http server: route already exist")
	}

	s.middlewares = append(s.middlewares, middlewares...)
}

func (s Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.ip, s.port), s.mux)
}

func chain(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
