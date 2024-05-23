package handlers

import (
	"net/http"

	"github.com/suyashmishra007/golang-web/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, "home.page.tmpl")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, "about.page.tmpl")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
