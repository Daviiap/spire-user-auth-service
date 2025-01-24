package infra

import (
	"user_auth_service/domain"
	repository "user_auth_service/domain/repository"
)

type tokenRepository struct {
	repository.TokenRepository
	tokens []domain.Token
}

func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{}
}

func (r *tokenRepository) Save(token domain.Token) error {
	r.tokens = append(r.tokens, token)

	return nil
}

func (r *tokenRepository) Get(tokenValue string) (domain.Token, error) {
	for _, token := range r.tokens {
		if token.Value == tokenValue {
			return token, nil
		}
	}

	return domain.Token{}, nil
}

func (r *tokenRepository) Delete(token domain.Token) error {
	for i, t := range r.tokens {
		if t.Value == token.Value {
			r.tokens = append(r.tokens[:i], r.tokens[i+1:]...)
			return nil
		}
	}

	return nil
}

func (r *tokenRepository) IsValid(token domain.Token) (bool, error) {
	token, err := r.Get(token.Value)

	if err != nil {
		return false, err
	}

	return token.IsValid, nil
}

func (r *tokenRepository) Update(token domain.Token) error {
	for i, savedToken := range r.tokens {
		if token.Value == savedToken.Value {
			r.tokens[i] = token
			return nil
		}
	}

	return nil
}
