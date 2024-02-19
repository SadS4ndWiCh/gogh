package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SadS4ndWiCh/gogh/http/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Failed to load env file")
	// }

	redisOpts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(redisOpts)

	userHandler := routes.NewUserHandler(redisClient)
	mux.HandleFunc("GET /users/{username}", userHandler.GetUser)
	mux.HandleFunc("GET /users/{username}/repos", userHandler.GetRepos)

	fmt.Println("🐙 Server is running at http://localhost:3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
