package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

type HtmlReturnResult struct {
	Documents []Document
	Drugs     []DrugInfo
}

type Document struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DrugInfo struct {
	Id                   string
	CyrillicName         string
	LatinName            string
	GenericName          string
	EANCode              string
	ATC                  string
	Form                 string
	Strength             string
	Packaging            string
	Content              string
	IssuanceMethod       string
	Warnings             string
	Manufacturer         string
	PlaceOfManufacturing string
	ApprovalHolder       string
	SolutionNumber       string
	SolutionDate         string
	ValidityDate         string
	RetailPrice          string
	WholesalePrice       string
	ReferencePrice       string
	FundPin              string
	UserGuide            string
	SummaryReport        string
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

		// The current Meilisearch Go package doesn't support MultiSearch so we need to hack it together until
		// we find a better way to do this (or just fork the package and fix it ourselves, but no time now)
		if query != "" {
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

			// Search drugs
			searchRes, err = client.Index("drug-registry").Search(query,
				&meilisearch.SearchRequest{
					Limit: 6,
				})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Extremely wasteful since JSON, work it out with a decoder
			results = searchRes.Hits
			var drugs []DrugInfo
			jsonString, _ = json.Marshal(results)
			json.Unmarshal(jsonString, &drugs)

			HtmlReturnResult := HtmlReturnResult{
				Documents: documents,
				Drugs:     drugs,
			}
			tmpl := template.Must(template.ParseFiles("views/searchResult.html"))
			tmpl.Execute(w, HtmlReturnResult)
		}
	})

	http.HandleFunc("/getDrugInfo", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.Form.Get("drugId")

		searchRes, err := client.Index("drug-registry").Search(query,
			&meilisearch.SearchRequest{
				Limit: 1,
			})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		results := searchRes.Hits
		var drug []DrugInfo
		jsonString, _ := json.Marshal(results)
		json.Unmarshal(jsonString, &drug)

		tmpl := template.Must(template.ParseFiles("views/drugModal.html"))
		tmpl.Execute(w, drug)

	})

	http.ListenAndServe(":8080", nil)
}
