package registry

import (
	"github.com/yona3/go-samples/clean_architecture/ent"
	"github.com/yona3/go-samples/clean_architecture/infrastructure/api/handler"
	database "github.com/yona3/go-samples/clean_architecture/infrastructure/database/repository"
	"github.com/yona3/go-samples/clean_architecture/interface/controllers"
	"github.com/yona3/go-samples/clean_architecture/interface/presenters"
	"github.com/yona3/go-samples/clean_architecture/usecase/presenter"
	"github.com/yona3/go-samples/clean_architecture/usecase/repository"
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
	return handler.NewAppHandler(i)
}

// user

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controllers.UserController {
	return controllers.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return database.NewUserRepository(i.conn)
}

func (i *interactor) NewUserPresenter() presenter.UserPresenter {
	return presenters.NewUserPresenter()
}

// test

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
