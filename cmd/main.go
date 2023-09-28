package main

import (
	"github.com/AminN77/we-connect/api/controller"
	"github.com/AminN77/we-connect/cmd/setup"
	"github.com/AminN77/we-connect/internal"
	"github.com/AminN77/we-connect/internal/agent"
	"github.com/AminN77/we-connect/pkg/csv"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// agent
	repo := internal.NewMongoRepository()
	//runAgent(repo)

	srv := internal.NewService(repo)
	con := controller.NewController(srv)

	// setup router
	router := setup.SetRouter(con)

	// run
	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		panic(err)
	}
}

func runAgent(repo internal.Repository) {
	file, err := os.OpenFile(os.Getenv("CSV_FILE_PATH"),
		os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	c := csv.New()
	ag := agent.New(repo, c)
	ag.Run(file)
}
