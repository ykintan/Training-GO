package postgresgorm

import (
	"context"
	"errors"
	"go-session16-crud-user-compose/entity"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db GormDBIface
}

// kalo ke db konstucture nya ke pgxpooliface yg sbelumnya sudah kita buat
func NewUserRepository(db GormDBIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{},
			err
	}
	return *user,
		nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{},
				nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.User{},
			err
	}
	return user,
		nil
}

func (r *userRepository) UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error) {
	var existinguser entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").First(&existinguser, id).Error; err != nil {
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.User{}, err
	}
	return existinguser, nil
}

func (r *userRepository) DeleteUserByID(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,
				nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
