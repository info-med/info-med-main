package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/types"
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

		// Search symposiums
		searchRes, err = meilisearchClient.Index("symposium-registry").Search(query,
			&meilisearch.SearchRequest{
				Limit: 6,
			})

		if err != nil {
			fmt.Println(err)
			panic("Error")
		}

		// Extremely wasteful since JSON, work it out with a decoder
		results = searchRes.Hits
		var symposiums []types.Symposium
		jsonString, _ = json.Marshal(results)
		json.Unmarshal(jsonString, &symposiums)

		// Search drugstores
		searchRes, err = meilisearchClient.Index("drugstore-registry").Search(query,
			&meilisearch.SearchRequest{
				Limit: 6,
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
			Symposiums: symposiums,
		}
		tmpl := template.Must(template.ParseFiles("views/searchResult.html"))
		tmpl.Execute(w, HtmlReturnResult)
	}
}

func HandleGetDrugstoreInfo(w http.ResponseWriter, r *http.Request, meilisearchClient *meilisearch.Client) {
	r.ParseForm()
	query := r.Form.Get("drugstoreId")

	searchRes, err := meilisearchClient.Index("drugstore-registry").Search(query,
		&meilisearch.SearchRequest{
			Limit: 1,
		})

	if err != nil {
		fmt.Println(err)
		panic("Error")
	}

	results := searchRes.Hits
	var drugstore []types.Drugstore
	jsonString, _ := json.Marshal(results)
	json.Unmarshal(jsonString, &drugstore)
	tmpl := template.Must(template.ParseFiles("views/drugstoreModal.html"))
	tmpl.Execute(w, drugstore)
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

func CreateSymposium(w http.ResponseWriter, r *http.Request, meilisearchClient *meilisearch.Client) {
	r.ParseForm()
	var symposium types.Symposium
	symposium.Id = uuid.NewString()
	symposium.Type = r.Form.Get("type")
	symposium.Name = r.Form.Get("name")
	symposium.Points = r.Form.Get("points")

	symposiumIndex := meilisearchClient.Index("symposium-registry")

	_, err := symposiumIndex.AddDocuments(symposium)

	if err != nil {
		fmt.Println(err)
		// TODO Return error message to UI via htmx
	}

	// TODO return success message to UI via htmx
}
