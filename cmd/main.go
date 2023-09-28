package main

import (
	"github.com/AminN77/we-connect/api/controller"
	"github.com/AminN77/we-connect/cmd/setup"
	"github.com/AminN77/we-connect/internal"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repo := internal.NewMongoRepository()
	srv := internal.NewService(repo)
	con := controller.NewController(srv)

	// setup router
	router := setup.SetRouter(con)

	// run
	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		panic(err)
	}
}
