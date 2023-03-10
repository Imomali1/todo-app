package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Imomali1/todo-app/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Configs, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
