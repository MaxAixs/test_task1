package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type SrvCfg struct {
	Port                string
	ReaderHeaderTimeOut time.Duration
	WriterHeaderTimeOut time.Duration
	IdleTimeout         time.Duration
}

func (s *Server) RunServer(cfg SrvCfg, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         cfg.Port,
		Handler:      handler,
		ReadTimeout:  cfg.ReaderHeaderTimeOut,
		WriteTimeout: cfg.WriterHeaderTimeOut,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(context.Background())
}
