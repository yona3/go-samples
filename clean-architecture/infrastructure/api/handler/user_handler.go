package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/interface/controllers"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
}

func NewUserHandler(c controllers.UserController) UserHandler {
	return &userHandler{c}
}

func (h *userHandler) Create(c *gin.Context) {
	req := &model.User{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Message: err.Error(),
		})
		return
	}

	if err := h.userController.Create(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Message{
		Message: "success",
	})
}

func (h *userHandler) GetAll(c *gin.Context) {
	req := []*model.User{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Message: err.Error(),
		})
		return
	}

	u, err := h.userController.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, u)
}
