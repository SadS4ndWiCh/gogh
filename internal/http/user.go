package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SadS4ndWiCh/gogh/internal/store"
	"github.com/SadS4ndWiCh/gogh/pkg/gh"
)

type UserHandler struct {
	cache store.Store
}

func NewUserHandler(cache store.Store) UserHandler {
	return UserHandler{cache: cache}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("user:%s", username)
	if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil {
		w.Write([]byte(cached.(string)))
		return
	}

	user, err := gh.GetUser(username)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := h.cache.Set(r.Context(), r.URL.Path, string(json)); err != nil {
		log.Printf("[CACHE] Failed to set data: %s\n", err)
	}

	w.Write(json)
}

func (h *UserHandler) GetRepos(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNumber = 1
	}

	cacheKey := fmt.Sprintf("user:repos(%d):%s", pageNumber, username)
	if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil {
		w.Write([]byte(cached.(string)))
		return
	}

	repositories, err := gh.GetRepositories(username, pageNumber)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(repositories)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := h.cache.Set(r.Context(), cacheKey, string(json)); err != nil {
		log.Printf("[CACHE] Failed to set data: %s\n", err)
	}

	w.Write(json)
}
