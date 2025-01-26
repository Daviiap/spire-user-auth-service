package usecases

import (
	"user_auth_service/domain/repository"
)

type ValidateTokenInput struct {
	Token string
}

type ValidateTokenOutput struct {
	Valid bool `json:"is_valid"`
}

type ValidateTokenUseCase struct {
	UseCase[ValidateTokenInput, ValidateTokenOutput]
	repository *repository.TokenRepository
}

func NewValidateTokenUseCase(repository *repository.TokenRepository) UseCase[ValidateTokenInput, ValidateTokenOutput] {
	return &ValidateTokenUseCase{repository: repository}
}

func (uc *ValidateTokenUseCase) Execute(input ValidateTokenInput) (ValidateTokenOutput, error) {
	token, err := (*uc.repository).Get(input.Token)

	if err != nil {
		return ValidateTokenOutput{Valid: false}, err
	}

	if token.IsValid {
		token.IsValid = false
		(*uc.repository).Update(token)

		return ValidateTokenOutput{Valid: true}, nil
	}
	return ValidateTokenOutput{Valid: false}, nil
}
