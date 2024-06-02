package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SadS4ndWiCh/gogh/internal/store"
)

type Server struct {
	addr  string
	cache store.Store
}

func NewServer(cacheStore store.Store) *http.Server {
	port := os.Getenv("PORT")
	srv := &Server{
		addr:  fmt.Sprintf(":%s", port),
		cache: cacheStore,
	}

	return &http.Server{
		Addr:    srv.addr,
		Handler: srv.Bootstrap(),
	}
}
