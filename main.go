package main

import (
	"bnlogic/data"
	"bnlogic/handlers"
	"bnlogic/middleware"

	"github.com/PaulDong/golaravel"
)

type application struct {
	App        *golaravel.Golaravel
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
