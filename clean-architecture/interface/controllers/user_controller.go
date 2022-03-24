package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	Create(ctx *gin.Context, u *model.User) error
	GetAll(ctx *gin.Context) ([]*model.User, error)
}

func NewUserController(s service.UserService) UserController {
	return &userController{s}
}

func (c *userController) Create(ctx *gin.Context, u *model.User) error {
	return c.userService.Create(ctx, u)
}

func (c *userController) GetAll(ctx *gin.Context) ([]*model.User, error) {
	users, err := c.userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
