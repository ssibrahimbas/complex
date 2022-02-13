package http

import (
	"context"
	"net/http"
)

type Http struct {
	server *http.Server
}

func NewServer(handler http.Handler, host string, port string) *Http {
	return &Http{
		server: &http.Server{
			Addr:    host + ":" + port,
			Handler: handler,
		},
	}
}

func (h *Http) Run() error {
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
