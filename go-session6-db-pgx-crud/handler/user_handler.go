package handler

import (
	"net/http"
	"strings"
	"training-go/go-session6-db-pgx-crud/entity"
	"training-go/go-session6-db-pgx-crud/service"

	"strconv"

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
		errMsg := err.Error()
		errMsg = fieldErrorMessage(errMsg)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	createdUser, err := h.userService.CreateUser(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, createdUser)

}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.userService.GetUserByID(ctx.Request.Context(), id)
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
	updatedUser, err := h.userService.UpdateUserByID(ctx.Request.Context(), id, user)
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
	err = h.userService.DeleteUserByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func fieldErrorMessage(errorMessage string) string {
	switch {
	case strings.Contains(errorMessage, "'name', failed on the 'required' tag"):
		return "Name is mandatory"
	case strings.Contains(errorMessage, "'name', failed on the 'min' tag"):
		return "Name must be at least 3 characters"
	case strings.Contains(errorMessage, "'email', failed on the 'required' tag"):
		return "Email is mandatory"
	case strings.Contains(errorMessage, "'email', failed on the 'email' tag"):
		return "Email is not valid"
	}
	return errorMessage
}
