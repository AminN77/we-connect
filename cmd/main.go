package main

import (
	"github.com/AminN77/we-connect/api/controller"
	"github.com/AminN77/we-connect/cmd/setup"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	con := controller.NewController()

	// setup router
	router := setup.SetRouter(con)

	// run
	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		panic(err)
	}
}
