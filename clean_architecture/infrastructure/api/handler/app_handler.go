package handler

type AppHandler interface {
	NewTestHandler() TestHandler
	NewUserHandler() UserHandler
}

type Interactor interface {
	AppHandler
}

type appHandler struct {
	interactor AppHandler
}

func NewAppHandler(i Interactor) AppHandler {
	return &appHandler{i}
}

func (ah *appHandler) NewTestHandler() TestHandler {
	return ah.interactor.NewTestHandler()
}

func (ah *appHandler) NewUserHandler() UserHandler {
	return ah.interactor.NewUserHandler()
}
