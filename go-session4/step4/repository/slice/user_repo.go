package slice

import "training-go/go-session4/step4/entity"

type IUserRepository interface {
	GetAllUsers() []entity.User
}

type userRepository struct {
	db     []entity.User //slice untuk menyimpan data user , bisa disebut array dlm bahasa lain
	nextID int
}

func NewUserRepository(db []entity.User) IUserRepository {
	return &userRepository{
		db:     db,
		nextID: 1,
	}
}

func (r *userRepository) GetAllUsers() []entity.User {
	return r.db
}
