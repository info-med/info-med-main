package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

type Document struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	fmt.Println("Started Server")
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
				Limit: 6,
			})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Extremely wasteful since JSON, work it out with a decoder
		results := searchRes.Hits
		var documents []Document
		jsonString, _ := json.Marshal(results)
		json.Unmarshal(jsonString, &documents)

		tmpl := template.Must(template.ParseFiles("views/searchResult.html"))
		tmpl.Execute(w, documents)
	})

	http.ListenAndServe(":8080", nil)
}
