package main

import (
	"fmt"
	"github.com/bndroll/go-project-01/pkg/config"
	"github.com/bndroll/go-project-01/pkg/handlers"
	"github.com/bndroll/go-project-01/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
