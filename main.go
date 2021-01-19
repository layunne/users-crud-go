package main

import (
	"log"

	"gitlab.com/layunne/users-crud-go/application"
	"gitlab.com/layunne/users-crud-go/config"
	"gitlab.com/layunne/users-crud-go/controller"
	"gitlab.com/layunne/users-crud-go/repository"
	"gitlab.com/layunne/users-crud-go/services"
)

func main() {

	env := config.NewEnv()

	usersRepository := repository.NewUsersRepository()

	rideService := services.NewUsersService(usersRepository)

	webController := controller.NewUsersWebController(rideService)

	webServer := application.NewUsersWebServer(env, webController)

	servers := []application.Server{webServer}

	exit := make(chan error)
	for _, server := range servers {
		go func(server application.Server) {
			exit <- server.Start()
		}(server)
	}

	log.Fatalln(<-exit)
}
