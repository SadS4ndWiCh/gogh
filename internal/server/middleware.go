package server

import (
	"log"
	"net/http"
)

type RouteFunc func(w http.ResponseWriter, r *http.Request)
type MiddlewareFunc func(RouteFunc) RouteFunc

func logMiddleware(fn RouteFunc) RouteFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)

		log.Printf("%s %s", r.Method, r.URL.Path)
	}
}
