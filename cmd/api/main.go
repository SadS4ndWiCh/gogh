package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SadS4ndWiCh/gogh/internal/server"
	"github.com/SadS4ndWiCh/gogh/internal/store"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cacheStore := store.NewRedisStore(os.Getenv("REDIS_URL"))
	srv := server.NewServer(cacheStore)

	fmt.Println("🐙 Server is running")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
