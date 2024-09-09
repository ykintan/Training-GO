package service

import (
	"fmt"
	"training-go/go-session5-validator/entity"
	"training-go/go-session5-validator/repository/slice"
)

type IUserService interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, error)
	UpdateUserByID(id int, user entity.User) (entity.User, error)
	DeleteUserByID(id int) error
	GetAllUsers() []entity.User
}

// rooter - handler - service - repository - db
type userService struct {
	userRepo slice.IUserRepository
}

func NewUserService(userRepo slice.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *entity.User) entity.User {
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (entity.User, error) {
	user, found := s.userRepo.GetUserByID(id)
	if !found {
		return entity.User{}, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (s *userService) UpdateUserByID(id int, user entity.User) (entity.User, error) {
	updatedUser, found := s.userRepo.UpdateUserByID(id, user)
	if !found {
		return entity.User{}, fmt.Errorf("user with id %d not found", id)
	}
	return updatedUser, nil
}

func (s *userService) DeleteUserByID(id int) error {
	if !s.userRepo.DeleteUserByID(id) {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}

func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
