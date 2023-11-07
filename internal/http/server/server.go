package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/marshevms/two_gis/internal/logger"
)

var mddlwrs []func(next http.Handler) http.Handler
var routeExist bool

type Server struct {
	srv  *http.Server
	mux  *http.ServeMux
	ip   string
	port int64
}

func New(ip string, port int64) *Server {
	mux := http.NewServeMux()

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", ip, port),
			Handler: mux,
		},

		ip:   ip,
		port: port,
		mux:  mux,
	}
}

func (s Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	routeExist = true
	s.mux.Handle(pattern, chain(http.HandlerFunc(handler), mddlwrs...))
}

func (s Server) Handle(pattern string, handler http.Handler) {
	routeExist = true
	s.mux.Handle(pattern, chain(handler, mddlwrs...))
}

// SetMiddleware устанавливает миддлвари, должно быть вызвано до установки роута
func (s Server) AddMiddleware(middlewares ...func(next http.Handler) http.Handler) {
	if routeExist {
		panic("http server: route already exist")
	}

	mddlwrs = append(mddlwrs, middlewares...)
}

func (s Server) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	logger.Infof("server is started: %s:%d", s.ip, s.port)

	return s.srv.Serve(l)
}

func (s Server) Stop(ctx context.Context) error {
	logger.Infof("server is stopping: %s:%d", s.ip, s.port)
	defer logger.Infof("server is stopped: %s:%d", s.ip, s.port)
	return s.srv.Shutdown(ctx)
}

func chain(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
