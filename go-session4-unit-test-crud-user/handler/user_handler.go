package handler

import (
	"net/http"
	"strconv"
	"training-go/go-session4-unit-test-crud-user/entity"
	"training-go/go-session4-unit-test-crud-user/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreatedUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}
type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreatedUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser := h.userService.CreateUser(&user)
	ctx.JSON(http.StatusOK, createdUser)

}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := h.userService.UpdateUserByID(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.userService.DeleteUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users := h.userService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}
