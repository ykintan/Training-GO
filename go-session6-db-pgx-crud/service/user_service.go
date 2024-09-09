package service

import (
	"context"
	"fmt"
	"training-go/go-session6-db-pgx-crud/entity"
	postgrespgx "training-go/go-session6-db-pgx-crud/repository/postgres_pgx"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// rooter - handler - service - repository - db
type userService struct {
	userRepo postgrespgx.IUserRepository
}

func NewUserService(userRepo postgrespgx.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error created user : %v", err)
	}
	return createdUser, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user not found : %v", err)
	}
	return user, nil
}

func (s *userService) UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := s.userRepo.UpdateUserByID(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed updated user : %v", err)
	}
	return updatedUser, nil
}

func (s *userService) DeleteUserByID(ctx context.Context, id int) error {
	err := s.userRepo.DeleteUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed deleted user : %v", err)
	}
	return nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get all users : %v", err)
	}
	return users, nil
}
