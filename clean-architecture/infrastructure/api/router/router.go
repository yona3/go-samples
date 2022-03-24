package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/infrastructure/api/handler"
)

func NewRouter(r *gin.Engine, h handler.AppHandler) {
	// test
	testHandler := h.NewTestHandler()
	r.GET("/test", testHandler.Test)

	// user
	userHandler := h.NewUserHandler()
	r.GET("/users", userHandler.GetAll)
	r.POST("/users", userHandler.Create)
}
