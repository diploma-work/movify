package infrastructure

import (
	"context"
	"fmt"
	"github.com/diploma/internal/config"
	"net/http"
)

type server struct {
	httpServer *http.Server
}

func NewServer(cfg config.ServerConfig, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Port),
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes << 20,
		},
	}
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
