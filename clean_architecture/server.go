package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean_architecture/infrastructure/api/router"
	"github.com/yona3/go-samples/clean_architecture/infrastructure/database"
	"github.com/yona3/go-samples/clean_architecture/registry"
)

func main() {
	// db
	conn := database.NewDB()

	// interactor
	i := registry.NewInteractor(conn)

	// gin
	r := gin.Default()

	// handler
	h := i.NewAppHandler()

	// router
	router.NewRouter(r, h)

	r.Run()
}
