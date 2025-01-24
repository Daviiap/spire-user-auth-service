package repository

import "user_auth_service/domain"

type TokenRepository interface {
	Save(domain.Token) error
	Get(string) (domain.Token, error)
	Delete(domain.Token) error
	IsValid(domain.Token) (bool, error)
	Update(domain.Token) error
}
