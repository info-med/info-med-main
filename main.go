package main

import (
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("views/index.html"))

		data := map[string]string{"TestKey": "TestValue"}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
