package presenters

import "github.com/yona3/go-samples/clean_architecture/domain/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (p *userPresenter) ResponseUsers(u []*model.User) []*model.User {
	return u
}
