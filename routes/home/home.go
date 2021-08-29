package home

import (
	"net/http"
	"templates"
	// "app"
	// "fmt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, []string{"home", "p"}, nil)
}
