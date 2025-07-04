package schemas

type ObjectSchema struct {
	Id                       string                  `json:"id"`
	AnvilId                  string                  `json:"anvilId"`
	Name                     string                  `json:"name"`
	PreviousNames            string                  `json:"previousNames"`
	SeriesId                 string                  `json:"seriesId"`
	CatalogueId              string                  `json:"catalogueId"`
	State                    string                  `json:"state"`
	ObjectClass              string                  `json:"objectClass"`
	Description              string                  `json:"description"`
	SizeBytes                int                     `json:"sizeBytes"`
	ApprovalStatus           ObjectApprovalsSchema   `json:"approvalStatus"`
	AttributeValues          []ObjectAttributeSchema `json:"attributeValues"`
	CreatedBy                string                  `json:"createdBy"`
	DateCreated              string                  `json:"dateCreated"`
	LegacyPipelineNotYetInS3 bool                    `json:"legacyPipelineNotYetInS3"`
}

type ObjectApprovalsSchema struct {
	ResolvedAnvilObjectProcessFlow string `json:"resolvedAnvilObjectProcessFlow"`
	AnvilObjectProcessFlowId       string `json:"anvilObjectProcessFlowId"`
	State                          string `json:"state"`
	StartedBy                      string `json:"startedBy"`
	DateStarted                    string `json:"dateStarted"`
	ApprovedBy                     string `json:"approvedBy"`
	DateApproved                   string `json:"dateApproved"`
	RejectedBy                     string `json:"rejectedBy"`
	DateRejected                   string `json:"dateRejected"`
	RejectedReason                 string `json:"rejectedReason"`
}

type ObjectAttributeSchema struct {
	AttributeName string `json:"attributeName"`
	Value         string `json:"value"`
}
