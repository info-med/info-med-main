package main

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/handlers"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Started Server")
	meilisearchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: os.Getenv("MEILISEARCH_MASTER_KEY"),
	})
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Handlers
	http.HandleFunc("/", handlers.HandleHomePage)
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSearch(w, r, meilisearchClient)
	})
	http.HandleFunc("/getDrugInfo", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGetDrugInfo(w, r, meilisearchClient)
	})
	http.HandleFunc("/getDrugstoreInfo", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGetDrugstoreInfo(w, r, meilisearchClient)
	})

	// Renderers
	http.HandleFunc("/renderIndex", handlers.RenderSearch)
	http.HandleFunc("/renderSymposium", handlers.RenderSymposium)
	http.HandleFunc("/renderAboutUs", handlers.RenderAboutUs)

	// CRD Routes
	http.HandleFunc("/createSymposium", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateSymposium(w, r, meilisearchClient)
	})

	http.ListenAndServe(":9990", nil)
}
