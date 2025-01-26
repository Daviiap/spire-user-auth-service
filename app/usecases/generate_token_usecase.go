package usecases

import (
	"crypto/rand"
	"encoding/hex"
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
	tokenValue, err := uc.generateRandomToken()
	if err != nil {
		return GenerateTokenOutput{}, err
	}

	token := domain.Token{
		Value:   tokenValue,
		IsValid: true,
	}

	(*uc.repository).Save(token)

	return GenerateTokenOutput{Token: token.Value}, nil
}

func (uc *GenerateTokenUseCase) generateRandomToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	return token, nil
}
