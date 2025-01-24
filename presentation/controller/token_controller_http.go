package controller

import (
	"encoding/json"
	"net/http"
	"user_auth_service/app/factory"
	"user_auth_service/app/usecases"
	"user_auth_service/presentation"
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

		if err := json.NewEncoder(w).Encode(token.Token); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}

	})
}
