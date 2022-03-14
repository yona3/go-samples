package registry

import (
	"github.com/yona3/go-samples/clean_architecture/ent"
	"github.com/yona3/go-samples/clean_architecture/infrastructure/api/handler"
	"github.com/yona3/go-samples/clean_architecture/interface/controllers"
	"github.com/yona3/go-samples/clean_architecture/interface/presenters"
	"github.com/yona3/go-samples/clean_architecture/usecase/presenter"
	"github.com/yona3/go-samples/clean_architecture/usecase/service"
)

type interactor struct {
	conn *ent.Client
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *ent.Client) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewTestHandler()
}

func (i *interactor) NewTestHandler() handler.TestHandler {
	return handler.NewTestHandler(i.NewTestController())
}

func (i *interactor) NewTestController() controllers.TestController {
	return controllers.NewTestController(i.NewTestService())
}

func (i *interactor) NewTestService() service.TestService {
	return service.NewTestService(i.NewTestPresenter())
}

func (i *interactor) NewTestPresenter() presenter.TestPresenter {
	return presenters.NewTestPresenter()
}
