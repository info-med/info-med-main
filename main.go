package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"html/template"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server")
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		data := map[string]string{"TestKey": "TestValue"}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.Form.Get("search")

		searchRes, err := client.Index("evidence-based-medicine").Search(query,
			&meilisearch.SearchRequest{
				Limit: 8,
			})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tmpl := template.Must(template.ParseFiles("views/searchResult.html"))
		tmpl.Execute(w, searchRes.Hits)
	})

	http.ListenAndServe(":8080", nil)
}
