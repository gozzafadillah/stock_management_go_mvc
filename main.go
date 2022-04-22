package main

import (
	"gozzafadillah/config"
	"gozzafadillah/routes"
)

func main() {
	config.Init()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
