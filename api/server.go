package api

import (
	"log"
	"net/http"

	"github.com/SadS4ndWiCh/gogh/http/routes"
	"github.com/SadS4ndWiCh/gogh/store"
)

type Server struct {
	addr  string
	cache store.Store
}

func NewServer(addr string, cacheStore store.Store) *Server {
	return &Server{
        addr:  addr,
		cache: cacheStore,
	}
}

func (s *Server) bootstrap() *http.ServeMux {
	api := http.NewServeMux()

	userHandler := routes.NewUserHandler(s.cache)
	api.HandleFunc("GET /users/{username}", logMiddleware(userHandler.GetUser))
	api.HandleFunc("GET /users/{username}/repos", logMiddleware(userHandler.GetRepos))

	return api
}

func (s *Server) Listen() error {
	log.Printf("🐙 Server is running at %s\n", s.addr)

	return http.ListenAndServe(s.addr, s.bootstrap())
}
