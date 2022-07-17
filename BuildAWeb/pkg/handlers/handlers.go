package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/render"
)

// Home is the handler for Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for About Page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
