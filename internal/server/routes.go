package server

import (
	"net/http"

	routes "github.com/SadS4ndWiCh/gogh/internal/http"
)

func (s *Server) Bootstrap() *http.ServeMux {
	api := http.NewServeMux()

	userHandler := routes.NewUserHandler(s.cache)
	api.HandleFunc("GET /users/{username}", logMiddleware(userHandler.GetUser))
	api.HandleFunc("GET /users/{username}/repos", logMiddleware(userHandler.GetRepos))

	return api
}
