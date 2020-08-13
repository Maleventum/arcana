// Package Arcana is our database service for storing and retrieving events, it
// uses the SQL driver that works for most databases, local tests are performed
// on SQLite, during deployment I expect to use a cloud service.
package main

import (
	"log"
	"net/http"

	_ "github.com/Maleventum/arcana/docs" // needed for the documentation

	"github.com/Maleventum/arcana/service/database"
	"github.com/Maleventum/arcana/usecase"

	"github.com/Maleventum/arcana/controller"
	"github.com/Maleventum/arcana/router"
)

func event() *controller.EventController {
	storage := database.New()
	usecase := usecase.NewEvent(storage)
	controller := controller.NewEvent(usecase)
	return controller
}

func main() {
	healthController := controller.New()
	eventController := event()

	r := router.Init(healthController, eventController)

	log.Fatal(http.ListenAndServe(":8081", r))
}
