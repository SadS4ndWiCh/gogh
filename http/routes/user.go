package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SadS4ndWiCh/gogh/pkg/gh"
	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	redis *redis.Client
}

func NewUserHandler(redis *redis.Client) UserHandler {
	return UserHandler{redis: redis}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cached, err := uh.redis.Get(r.Context(), r.URL.Path).Result()
	if err == nil {
		w.Write([]byte(cached))
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

	if err := uh.redis.Set(r.Context(), r.URL.Path, string(json), time.Duration(2*time.Hour)).Err(); err != nil {
		log.Printf("[REDIS] Failed to set data: %s\n", err)
	}

	w.Write(json)
}

func (uh *UserHandler) GetRepos(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNumber = 1
	}

	cacheKey := fmt.Sprintf("%s/page=%d", r.URL.Path, pageNumber)
	cached, err := uh.redis.Get(r.Context(), cacheKey).Result()
	if err == nil {
		w.Write([]byte(cached))
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

	if err := uh.redis.Set(r.Context(), cacheKey, string(json), time.Duration(2*time.Hour)).Err(); err != nil {
		log.Printf("[REDIS] Failed to set data: %s\n", err)
	}

	w.Write(json)
}
