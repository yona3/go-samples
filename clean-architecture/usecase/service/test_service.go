package service

import (
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/usecase/presenter"
)

type testService struct {
	TestPresenter presenter.TestPresenter
}

type TestService interface {
	Test() ([]*model.Message, error)
}

func NewTestService(p presenter.TestPresenter) TestService {
	return &testService{p}
}

func (s *testService) Test() ([]*model.Message, error) {
	m := []*model.Message{
		{
			Message: "test",
		},
	}
	return s.TestPresenter.ResponseTest(m), nil
}
