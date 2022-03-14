package presenter

import "github.com/yona3/go-samples/clean_architecture/domain/model"

type TestPresenter interface {
	ResponseTest([]*model.Message) []*model.Message
}
