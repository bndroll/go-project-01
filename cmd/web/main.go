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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
