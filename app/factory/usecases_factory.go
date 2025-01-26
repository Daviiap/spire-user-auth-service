package factory

import (
	"user_auth_service/app/usecases"
	"user_auth_service/domain/repository"
)

type UseCasesFactory struct {
	tokenRepository *repository.TokenRepository
	userRepository  *repository.UserRepository
}

func NewUseCasesFactory(tokenRepository *repository.TokenRepository, userRepository *repository.UserRepository) *UseCasesFactory {
	return &UseCasesFactory{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
	}
}

func (f *UseCasesFactory) NewGenerateTokenUseCase() usecases.UseCase[usecases.GenerateTokenInput, usecases.GenerateTokenOutput] {
	return usecases.NewGenerateTokenUseCase(f.tokenRepository)
}

func (f *UseCasesFactory) NewValidateTokenUseCase() usecases.UseCase[usecases.ValidateTokenInput, usecases.ValidateTokenOutput] {
	return usecases.NewValidateTokenUseCase(f.tokenRepository)
}

func (f *UseCasesFactory) NewLoginUseCase() usecases.UseCase[usecases.LoginInput, usecases.LoginOutput] {
	return usecases.NewLoginUseCase(f.userRepository)
}
