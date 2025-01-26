package usecases

type LoginInput struct {
	User     string
	Password string
}

type LoginOutput struct {
	JWT string `json:"jwt"`
}

type LoginUseCase struct {
	UseCase[LoginInput, LoginOutput]
}

func NewLoginUseCase() UseCase[LoginInput, LoginOutput] {
	return &LoginUseCase{}
}

func (uc *LoginUseCase) Execute(input LoginInput) (LoginOutput, error) {
	return LoginOutput{JWT: "TOKEN"}, nil
}
