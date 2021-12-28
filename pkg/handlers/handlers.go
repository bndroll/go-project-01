package handlers

import (
	"github.com/bndroll/go-project-01/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
