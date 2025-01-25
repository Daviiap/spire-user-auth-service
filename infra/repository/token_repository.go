package infra

import (
	"sync"
	"user_auth_service/domain"
	repository "user_auth_service/domain/repository"
)

type tokenRepository struct {
	repository.TokenRepository
	mutex  sync.Mutex
	tokens []domain.Token
}

func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{
		tokens: []domain.Token{},
	}
}

func (r *tokenRepository) Save(token domain.Token) error {
	r.mutex.Lock()
	r.tokens = append(r.tokens, token)
	r.mutex.Unlock()
	return nil
}

func (r *tokenRepository) Get(tokenValue string) (domain.Token, error) {
	r.mutex.Lock()
	for _, token := range r.tokens {
		if token.Value == tokenValue {
			r.mutex.Unlock()
			return token, nil
		}
	}

	r.mutex.Unlock()
	return domain.Token{}, nil
}

func (r *tokenRepository) Delete(token domain.Token) error {
	r.mutex.Lock()
	for i, t := range r.tokens {
		if t.Value == token.Value {
			r.tokens = append(r.tokens[:i], r.tokens[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	r.mutex.Unlock()
	return nil
}

func (r *tokenRepository) Update(token domain.Token) error {
	r.mutex.Lock()
	for i, savedToken := range r.tokens {
		if token.Value == savedToken.Value {
			r.tokens[i] = token
			r.mutex.Unlock()
			return nil
		}
	}
	r.mutex.Unlock()
	return nil
}
