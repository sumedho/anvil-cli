package schemas

type OwnerSchema struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	GivenName string `json:"givenName"`
	LastName  string `json:"lastName"`
}

type CatalogueSchema struct {
	Id                                 string      `json:"id"`
	Name                               string      `json:"name"`
	Description                        string      `json:"description"`
	Version                            int         `json:"version"`
	IsOfficial                         bool        `json:"isOfficial"`
	IsDraft                            bool        `json:"isDraft"`
	Owner                              OwnerSchema `json:"owner"`
	OwnerUserName                      string      `json:"ownerUsername"`
	CreatedBy                          string      `json:"createdBy"`
	DateCreated                        string      `json:"dateCreated"`
	Prefix                             string      `json:"prefix"`
	CustomSeriesNaming                 string      `json:"customSeriesNaming"`
	CustomObjectNaming                 string      `json:"customObjectNaming"`
	ApprovalProcessProcessFLowSeriesId string      `json:"approvalProcessProcessFlowSeriesId"`
	ObjectClasses                      []string    `json:"objectClasses"`
}

type CatalogueSummarySchema struct {
	Limit              int               `json:"limit"`
	Offset             int               `json:"offset"`
	CatalogueSummaries []CatalogueSchema `json:"catalogueSummaries"`
	TotalCount         int               `json:"totalCount"`
}
