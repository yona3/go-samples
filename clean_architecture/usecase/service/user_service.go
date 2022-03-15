package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean_architecture/domain/model"
	"github.com/yona3/go-samples/clean_architecture/usecase/presenter"
	"github.com/yona3/go-samples/clean_architecture/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Create(ctx *gin.Context, u *model.User) error
	GetAll(ctx *gin.Context) ([]*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

func (s *userService) Create(ctx *gin.Context, u *model.User) error {
	return s.UserRepository.Store(ctx, u)
}

func (s *userService) GetAll(ctx *gin.Context) ([]*model.User, error) {
	users, err := s.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return s.UserPresenter.ResponseUsers(users), nil
}
