package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//setup service
	var mockUserDBInSlice []User
	userRepo := NewUserRepository(mockUserDBInSlice)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	//setup router
	SetupRouter(r, userHandler)

	log.Println("Server started at localhost:8080")
	r.Run("localhost:8080")

}

// ----------------------------------------------
type IUserRepository interface {
	GetAllUsers() []User
}

type userRepository struct {
	db     []User //slice untuk menyimpan data user , bisa disebut array dlm bahasa lain
	nextID int
}

func NewUserRepository(db []User) IUserRepository {
	return &userRepository{
		db:     db,
		nextID: 1,
	}
}

func (r *userRepository) GetAllUsers() []User {
	return r.db
}

// ----------------------------------------------
type IUserService interface {
	GetAllUsers() []User
}

// rooter - handler - service - repository - db
type userService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUsers() []User {
	return s.userRepo.GetAllUsers()
}

//----------------------------------------------

type IUserHandler interface {
	GetAllUsers(ctx *gin.Context)
}
type UserHandler struct {
	userService IUserService
}

func NewUserHandler(userService IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users := h.userService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func SetupRouter(r *gin.Engine, userHandler IUserHandler) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
}
