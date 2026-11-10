package services

import (
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
)

type UserServiceInterface interface {
	GetUserByID(id interface{}) (*models.User, error)
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id interface{}) (*models.User, error) {
	user := s.userRepo.GetById(id)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}
