package main

import (
	"go_microservices_skeleton/config"
	logger "go_microservices_skeleton/middleware"
	"go_microservices_skeleton/model"
	"go_microservices_skeleton/repository"
	"go_microservices_skeleton/server"
	"go_microservices_skeleton/usecase"

	log "github.com/sirupsen/logrus"
)

func init() {
	config.Load()
	logger.Setup()
}

func main() {
	conf := config.Db()
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := model.CreateConnection(conf)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repository := &repository.UserRepository{Db: db}
	userService := &usecase.Service{Repo: repository}

	// Run the server
	if err := server.StartGRPCServer(userService); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
