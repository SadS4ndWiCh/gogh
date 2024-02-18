package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SadS4ndWiCh/gogh/pkg/gh"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := gh.GetUser(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func GetRepos(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var pageNumber int
	n, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNumber = 1
	} else {
		pageNumber = n
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

	w.Write(json)
}
