package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer(cfg *Config, router *mux.Router) *http.Server {
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		Handler:      router,
	}

	return srv
}
