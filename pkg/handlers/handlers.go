package handlers

import (
	"github.com/bndroll/go-project-01/pkg/config"
	"github.com/bndroll/go-project-01/pkg/models"
	"github.com/bndroll/go-project-01/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateDate{})
}

func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World !"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateDate{
		StringMap: stringMap,
	})
}
