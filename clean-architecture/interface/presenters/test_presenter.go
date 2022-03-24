package presenters

import "github.com/yona3/go-samples/clean-architecture/domain/model"

type testPresenter struct {
}

type TestPresenter interface {
	ResponseTest(m []*model.Message) []*model.Message
}

func NewTestPresenter() TestPresenter {
	return &testPresenter{}
}

func (p *testPresenter) ResponseTest(m []*model.Message) []*model.Message {
	return m
}
