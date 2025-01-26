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

type TokenControllerHttp struct {
	HttpController
	httpServer      *presentation.HttpServer
	useCasesFactory *factory.UseCasesFactory
}

func NewTokenControllerHttp(httpServer *presentation.HttpServer, useCasesFactory *factory.UseCasesFactory) HttpController {
	return &TokenControllerHttp{
		httpServer:      httpServer,
		useCasesFactory: useCasesFactory,
	}
}

func (c *TokenControllerHttp) SetAllRoutes() {
	(*c.httpServer).AddRoute("/generate-token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token, _ := c.useCasesFactory.NewGenerateTokenUseCase().Execute(usecases.GenerateTokenInput{})

		if err := json.NewEncoder(w).Encode(token); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})

	(*c.httpServer).AddRoute("/validate-token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var request dto.ValidateTokenInput

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		output, err := c.useCasesFactory.NewValidateTokenUseCase().Execute(usecases.ValidateTokenInput{Token: request.Token})

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
