package simpleRestAPI

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServe *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServe = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServe.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServe.Shutdown(ctx)
}
