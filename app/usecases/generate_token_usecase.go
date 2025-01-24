package usecases

import (
	"math/rand"
	"time"
	"user_auth_service/domain"
	"user_auth_service/domain/repository"
)

type GenerateTokenInput struct{}

type GenerateTokenOutput struct {
	Token string `json:"token"`
}

type GenerateTokenUseCase struct {
	UseCase[GenerateTokenInput, GenerateTokenOutput]
	repository *repository.TokenRepository
}

func NewGenerateTokenUseCase(repository *repository.TokenRepository) UseCase[GenerateTokenInput, GenerateTokenOutput] {
	return &GenerateTokenUseCase{repository: repository}
}

func (uc *GenerateTokenUseCase) Execute(input GenerateTokenInput) (GenerateTokenOutput, error) {
	tokenValue := generateRandomToken(32)
	token := domain.Token{
		Value:   tokenValue,
		IsValid: true,
	}

	(*uc.repository).Save(token)

	return GenerateTokenOutput{Token: token.Value}, nil
}

func generateRandomToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, length)
	for i := range token {
		token[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(token)
}
