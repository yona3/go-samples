package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean_architecture/infrastructure/api/handler"
)

func NewRouter(r *gin.Engine, handler handler.AppHandler) {
	r.GET("/test", handler.Test)
}
