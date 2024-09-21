package postgrespgx

import (
	"context"
	"go-session16-crud-user-compose/entity"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxPoolIface interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	//db     []entity.User //slice untuk menyimpan data user , bisa disebut array dlm bahasa lain
	//nextID int
	db PgxPoolIface
}

// kalo ke db konstucture nya ke pgxpooliface yg sbelumnya sudah kita buat
func NewUserRepository(db PgxPoolIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) returning id"
	var id int
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return entity.User{}, err

	}
	user.ID = id
	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUserByID(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, id)
	if err != nil {
		log.Printf("error updating user: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		log.Printf("error deleting user: %v\n", err)
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("error getting all users: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("error scanning user: %v\n", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil

}
