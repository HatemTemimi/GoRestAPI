package service

import (
	"apigo/internal/user/models"
	"apigo/internal/user/repository"
	"errors"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func MakeUserService(repo repository.UserRepository) UserService {
	return UserService{UserRepository: repo}
}

func (u *UserService) Create(usr models.User) (*models.User, error) {
	user, err := u.UserRepository.Create(usr)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) FindByEmail(email string) (*models.User, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}
