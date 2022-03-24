package controllers

import (
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/usecase/service"
)

type testController struct {
	testService service.TestService
}

type TestController interface {
	Test() ([]*model.Message, error)
}

func NewTestController(s service.TestService) TestController {
	return &testController{s}
}

func (c *testController) Test() ([]*model.Message, error) {
	d, err := c.testService.Test()
	if err != nil {
		return nil, err
	}
	return d, err
}
