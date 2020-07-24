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
