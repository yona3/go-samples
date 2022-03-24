package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/infrastructure/api/router"
	"github.com/yona3/go-samples/clean-architecture/infrastructure/database"
	"github.com/yona3/go-samples/clean-architecture/registry"
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

	// close db
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(fmt.Sprintf("failed to close db: %v", err))
		}
	}()

	r.Run()
}
