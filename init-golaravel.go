package main

import (
	"bnlogic/data"
	"bnlogic/handlers"
	"bnlogic/middleware"
	"log"
	"os"

	"github.com/PaulDong/golaravel"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// init golaravel
	gol := &golaravel.Golaravel{}
	err = gol.New(path)
	if err != nil {
		log.Fatal(err)
	}
	gol.AppName = "bnlogic"
	myMiddleware := &middleware.Middleware{
		App: gol,
	}

	myHandlers := &handlers.Handlers{
		App: gol,
	}
	app := &application{
		App:        gol,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
