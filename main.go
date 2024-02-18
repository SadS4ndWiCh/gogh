package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SadS4ndWiCh/gogh/http/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{username}", routes.GetUser)

	fmt.Println("🐙 Server is running at http://localhost:3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
