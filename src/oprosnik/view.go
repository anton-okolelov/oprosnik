package oprosnik

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("resources/templates/*.html"))

func render(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
