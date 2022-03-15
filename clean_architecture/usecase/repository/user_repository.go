package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean_architecture/domain/model"
)

type UserRepository interface {
	Store(ctx *gin.Context, u *model.User) error
	FindAll(ctx *gin.Context) ([]*model.User, error)
}
