package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/types"
	"html/template"
	"net/http"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	// Check the docs for excecuting without data pls
	data := map[string]string{}
	tmpl.Execute(w, data)
}

func HandleSearch(w http.ResponseWriter, r *http.Request, meilisearchClient *meilisearch.Client) {
	r.ParseForm()
	query := r.Form.Get("search")

	// The current Meilisearch Go package doesn't support MultiSearch so we need to hack it together until
	// we find a better way to do this (or just fork the package and fix it ourselves, but no time now)
	if query != "" {
		searchRes, err := meilisearchClient.Index("evidence-based-medicine").Search(query,
			&meilisearch.SearchRequest{
				Limit: 6,
			})

		if err != nil {
			fmt.Println(err)
			panic("Error")
		}

		// Extremely wasteful since JSON, work it out with a decoder
		results := searchRes.Hits
		var documents []types.Document
		jsonString, _ := json.Marshal(results)
		json.Unmarshal(jsonString, &documents)

		// Search drugs
		searchRes, err = meilisearchClient.Index("drug-registry").Search(query,
			&meilisearch.SearchRequest{
				Limit: 6,
			})

		if err != nil {
			fmt.Println(err)
			panic("Error")
		}

		// Extremely wasteful since JSON, work it out with a decoder
		results = searchRes.Hits
		var drugs []types.DrugInfo
		jsonString, _ = json.Marshal(results)
		json.Unmarshal(jsonString, &drugs)

		// Search drugstores
		searchRes, err = meilisearchClient.Index("drugstore-registry").Search(query,
			&meilisearch.SearchRequest{
				Limit: 2,
			})

		if err != nil {
			fmt.Println(err)
			panic("Error")
		}

		// Extremely wasteful since JSON, work it out with a decoder
		results = searchRes.Hits
		var drugstores []types.Drugstore
		jsonString, _ = json.Marshal(results)
		json.Unmarshal(jsonString, &drugstores)

		HtmlReturnResult := types.HtmlReturnResult{
			Documents:  documents,
			Drugs:      drugs,
			Drugstores: drugstores,
		}
		tmpl := template.Must(template.ParseFiles("views/searchResult.html"))
		tmpl.Execute(w, HtmlReturnResult)
	}
}

func HandleGetDrugInfo(w http.ResponseWriter, r *http.Request, meilisearchClient *meilisearch.Client) {
	r.ParseForm()
	query := r.Form.Get("drugId")

	searchRes, err := meilisearchClient.Index("drug-registry").Search(query,
		&meilisearch.SearchRequest{
			Limit: 1,
		})

	if err != nil {
		fmt.Println(err)
		panic("Error")
	}

	results := searchRes.Hits
	var drug []types.DrugInfo
	jsonString, _ := json.Marshal(results)
	json.Unmarshal(jsonString, &drug)
	tmpl := template.Must(template.ParseFiles("views/drugModal.html"))
	tmpl.Execute(w, drug)

}
