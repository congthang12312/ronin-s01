package runner

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pkgerrors "github.com/pkg/errors"
)

// NewServer initializes and returns an instance of HTTP Server
func NewServer(handler http.Handler, addr string) *Server {
	s := Server{
		srv: &http.Server{
			Addr:         ":3000",
			Handler:      handler,
			ReadTimeout:  time.Minute,
			WriteTimeout: time.Minute,
		},
		shutdownGrace: 10 * time.Second,
	}
	s.srv.Addr = fmt.Sprintf(":%s", addr)
	return &s
}

// Server represents an HTTP server
type Server struct {
	srv           *http.Server
	shutdownGrace time.Duration
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	startupError := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		startupError <- s.srv.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-startupError:
		if err != http.ErrServerClosed { // ListenAndServe will always return a non-nil error
			return pkgerrors.WithStack(fmt.Errorf("startup failed: %w", err))
		}
		return nil
	case <-ctx.Done():
		return s.Stop()
	}
}

// Stop stops the HTTP server
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownGrace)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		if err = s.srv.Close(); err != nil {
			return pkgerrors.WithStack(fmt.Errorf("force shutdown failed: %w", err))
		}
	}

	return nil
}
