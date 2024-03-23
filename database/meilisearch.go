package database

import (
	"encoding/json"
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/types"
)

// The current Meilisearch Go package doesn't support MultiSearch so we need to hack it together until
// we find a better way to do this (or just fork the package and fix it ourselves, but no time now)
func Search(query string) types.HtmlReturnResult {
	searchRes, err := meilisearchClient.Index("evidence-based-medicine").Search(query,
		&meilisearch.SearchRequest{
			Limit:                6,
			AttributesToSearchOn: []string{"name", "url"},
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
      Limit: 1000,
			AttributesToSearchOn: []string{"CyrillicName", "LatinName", "Generic", "EANCode", "Content", "Manufacturer", "SolutionNumber", "FundPin"},
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
			Limit:                20,
			AttributesToSearchOn: []string{"Name", "Address", "Municipality"},
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

	// Search MKB10 Entries
	/* searchRes, err = meilisearchClient.Index("temp-mkb-registry").Search(query,
		&meilisearch.SearchRequest{
			Limit:                50,
			AttributesToSearchOn: []string{"ninja_column_2", "ninja_column_3"},
		})

	if err != nil {
		fmt.Println(err)
		panic("Error")
	}

	// Extremely wasteful since JSON, work it out with a decoder
	results = searchRes.Hits
	var mkbEntries []types.MKBEntry
	jsonString, _ = json.Marshal(results)
	json.Unmarshal(jsonString, &mkbEntries) */
  var mkbEntries []types.MKBEntry

	HtmlReturnResult := types.HtmlReturnResult{
		Documents:  documents,
		Drugs:      drugs,
		Drugstores: drugstores,
		MkbEntries: mkbEntries,
	}

	return HtmlReturnResult
}

func GetDrugInfo(query string) []types.DrugInfo {
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

	return drug
}

func GetDrugstoreInfo(query string) []types.Drugstore {
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

	return drugstore
}
