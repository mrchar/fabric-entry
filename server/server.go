package server

import (
	"context"
	"net/http"
)

type Server struct {
	server *http.Server
	router http.Handler
}

func New(addr string) *Server {
	router := newRouter()
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router: router,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
