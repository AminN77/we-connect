package main

import (
	"github.com/AminN77/we-connect/api/controller"
	"github.com/AminN77/we-connect/cmd/setup"
	"github.com/AminN77/we-connect/internal"
	"github.com/AminN77/we-connect/internal/agent"
	csvPkg "github.com/AminN77/we-connect/pkg/csv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// agent
	repo := internal.NewMongoRepository()
	runAgent(repo)

	// service & controller
	srv := internal.NewService(repo)
	con := controller.NewController(srv)

	// setup router
	router := setup.SetRouter(con)

	// run
	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		panic(err)
	}
}

// runAgent provides the src for loading csv data and runs the agent.
func runAgent(repo internal.Repository) {
	if v, err := strconv.ParseBool(os.Getenv("LOAD_DATA")); err != nil || v == false {
		log.Println("value for LOAD_DATA env is not valid or it is false. skipped.")
		return
	}

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

	c := csvPkg.New()
	ag := agent.New(repo, c)
	ag.Run(file)

	// gc will be called before the server goes up which will reduce the amount of CPU throttling
	// caused by stopped the world procedure after the server runs.
	runtime.GC()
}
