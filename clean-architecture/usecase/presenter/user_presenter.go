package presenter

import "github.com/yona3/go-samples/clean-architecture/domain/model"

// TODO: add response user

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
