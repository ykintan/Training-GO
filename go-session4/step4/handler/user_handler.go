package handler

import (
	"net/http"
	"training-go/go-session4/step4/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
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

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users := h.userService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}
