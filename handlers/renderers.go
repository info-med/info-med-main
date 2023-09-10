package handlers

import (
	"html/template"
	"net/http"
)

func RenderSearch(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/search.html"))

	// Check the docs for excecuting without data pls
	data := map[string]string{}
	tmpl.Execute(w, data)

}

func RenderSymposium(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/symposium.html"))

	// Check the docs for excecuting without data pls
	data := map[string]string{}
	tmpl.Execute(w, data)
}

func RenderAbout(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/about.html"))

	// Check the docs for excecuting without data pls
	data := map[string]string{}
	tmpl.Execute(w, data)
}
