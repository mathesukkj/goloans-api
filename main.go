package main

import (
	"log"

	"github.com/mathesukkj/goloans-api/routes"
)

func main() {
	app := routes.NewRouter()

	log.Fatal(app.Listen(":8080"))
}
