package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"user_auth_service/app/factory"
	"user_auth_service/app/usecases"
	"user_auth_service/presentation"
	"user_auth_service/presentation/dto"
)

type LoginControllerHttp struct {
	HttpController
	httpServer      *presentation.HttpServer
	useCasesFactory *factory.UseCasesFactory
}

func NewLoginControllerHttp(httpServer *presentation.HttpServer, useCasesFactory *factory.UseCasesFactory) HttpController {
	return &LoginControllerHttp{
		httpServer:      httpServer,
		useCasesFactory: useCasesFactory,
	}
}

func (c *LoginControllerHttp) SetAllRoutes() {
	(*c.httpServer).AddRoute("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var request dto.LoginInput

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		output, err := c.useCasesFactory.NewLoginUseCase().Execute(usecases.LoginInput{
			User:     request.User,
			Password: request.Password,
		})

		if err != nil {
			log.Printf("Error validating token: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})
}
