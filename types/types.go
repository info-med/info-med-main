package types

type HtmlReturnResult struct {
	Documents  []Document
	Drugs      []DrugInfo
	Drugstores []Drugstore
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Drugstore struct {
	Id           string
	Name         string
	Address      string
	Municipality string
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
