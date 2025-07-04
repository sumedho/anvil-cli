package schemas

type PrefixObjectSchema struct {
	Id           string `json:"id"`
	Prefix       string `json:"prefix"`
	Label        string `json:"label"`
	Description  string `json:"description"`
	CreatedBy    string `json:"createdBy"`
	DateCreated  string `json:"dateCreated"`
	ModifiedBy   string `json:"modifiedBy"`
	LastModified string `json:"lastModified"`
}

type PrefixSchema struct {
	CataloguePrefixes []PrefixObjectSchema `json:"cataloguePrefixes"`
	TotalCount        int                  `json:"totalCount"`
}
