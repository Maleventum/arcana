// Package Arcana is our database service for storing and retrieving events, it
// uses the SQL driver that works for most databases, local tests are performed
// on SQLite, during deployment I expect to use a cloud service.
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Maleventum/arcana/docs" // needed for the documentation

	"github.com/Maleventum/arcana/controller"
	"github.com/Maleventum/arcana/router"
)

func main() {
	fmt.Println("hello world")

	health := controller.New()
	r := router.Init(health)

	log.Fatal(http.ListenAndServe(":8081", r))
}
