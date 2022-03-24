package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/interface/controllers"
)

type testHandler struct {
	testController controllers.TestController
}

type TestHandler interface {
	Test(c *gin.Context)
}

func NewTestHandler(c controllers.TestController) TestHandler {
	return &testHandler{c}
}

func (h *testHandler) Test(c *gin.Context) {
	d, err := h.testController.Test()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, d)
}
