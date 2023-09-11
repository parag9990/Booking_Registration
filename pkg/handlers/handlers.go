package handlers

import (
	"net/http"

	"BookingsWebsite/pkg/config"
	"BookingsWebsite/pkg/models"
	"BookingsWebsite/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository represents a repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NeWHandlers sets the repository for the Handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// This function handles requests
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is my home page")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is my home page")
	stringmap := make(map[string]string)
	stringmap["test"] = "Hello Parag"
	remote_IP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remote_IP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	})

}
