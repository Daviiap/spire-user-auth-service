package usecases

import (
	"errors"
	"user_auth_service/domain/repository"
	"user_auth_service/utils"
)

type LoginInput struct {
	User     string
	Password string
}

type LoginOutput struct {
	JWT string `json:"jwt"`
}

type LoginUseCase struct {
	UseCase[LoginInput, LoginOutput]
	repository *repository.UserRepository
}

func NewLoginUseCase(repository *repository.UserRepository) UseCase[LoginInput, LoginOutput] {
	return &LoginUseCase{
		repository: repository,
	}
}

func (uc *LoginUseCase) Execute(input LoginInput) (LoginOutput, error) {
	user, err := (*uc.repository).GetByName(input.User)
	if err != nil {
		return LoginOutput{}, err
	}

	if !user.IsValidPassword(input.Password) {
		return LoginOutput{}, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(utils.UserInfo{
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		Organization: user.GetOrganization(),
	})
	if err != nil {
		return LoginOutput{}, err
	}

	return LoginOutput{JWT: token}, nil
}
