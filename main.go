package main

import (
	"log"
	"os"
	"user_auth_service/app/factory"
	infra "user_auth_service/infra/http"
	repository_impl "user_auth_service/infra/repository"
	"user_auth_service/presentation/controller"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf("No .env file found, proceeding with system environment variables.")
	}

	tokenRepository := repository_impl.NewTokenRepository()
	useCaseFactory := factory.NewUseCasesFactory(&tokenRepository)

	port := os.Getenv("SERVER_PORT")
	httpServer := infra.NewServer(port)
	tokenHttpController := controller.NewTokenControllerHttp(&httpServer, useCaseFactory)

	tokenHttpController.SetAllRoutes()

	log.Printf("Starting server on port %s", port)
	httpServer.Start()
}
