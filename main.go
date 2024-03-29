package main

import (
	"flag"
	"log"
	"os"

	"github.com/SadS4ndWiCh/gogh/api"
	"github.com/SadS4ndWiCh/gogh/store"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
    addr := flag.String("addr", ":3000", "The server address")
    flag.Parse()

    cacheStore := store.NewRedisStore(os.Getenv("REDIS_URL"))
    server := api.NewServer(*addr, cacheStore)

	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
