package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/iMeisa/errortrace"
	"github.com/iMeisa/weed/internal/config"
	"github.com/iMeisa/weed/internal/models"
	"github.com/iMeisa/weed/internal/render"
	"github.com/iMeisa/weed/internal/repository"
	"log"
	"net/http"
)

// Main handlers file

// Repo the repository used by the handlers
var Repo *Repository

// Template folder consts
const (
	PublicDir = "public"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Public handler directs you to any public page based on the url
func (m *Repository) Public(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	templatePage := fmt.Sprintf("%s.page.tmpl", page)
	render.Template(w, r, PublicDir, templatePage, &models.TemplateData{})
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, PublicDir, "home.page.tmpl", &models.TemplateData{})
}

func writeResp(w http.ResponseWriter, resp models.JsonResponse) {
	respJSON, err := json.Marshal(resp)
	if err != nil {
		trace := errortrace.NewTrace(err)
		trace.Read()
		log.Println("Error marshalling response struct to json")
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(respJSON)
}
