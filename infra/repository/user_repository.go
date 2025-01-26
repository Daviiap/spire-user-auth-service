package infra

import (
	"user_auth_service/domain"
	repository "user_auth_service/domain/repository"
)

type userRepository struct {
	repository.UserRepository
	users []domain.User
}

func NewUserRepository() repository.UserRepository {
	user := domain.NewUser("jonh_doe", "john.doe@example.com", "ACME Inc.", "password")
	return &userRepository{
		users: []domain.User{
			*user,
		},
	}
}

func (r *userRepository) Save(user domain.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *userRepository) Get(userID string) (domain.User, error) {
	for _, user := range r.users {
		if user.GetID() == userID {
			return user, nil
		}
	}
	return domain.User{}, nil
}

func (r *userRepository) GetByName(userName string) (domain.User, error) {
	for _, user := range r.users {
		if user.GetName() == userName {
			return user, nil
		}
	}
	return domain.User{}, nil
}

func (r *userRepository) Delete(user domain.User) error {
	for i, u := range r.users {
		if u.GetID() == user.GetID() {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *userRepository) Update(user domain.User) error {
	for i, u := range r.users {
		if u.GetID() == user.GetID() {
			r.users[i] = user
			return nil
		}
	}
	return nil
}
