package service

import (
	"training-go/go-session4/step4/entity"
	"training-go/go-session4/step4/repository/slice"
)

type IUserService interface {
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

func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
