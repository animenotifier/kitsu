package kitsu

// Mapping ...
type Mapping struct {
	ID string `json:"id"`
	// Type string `json:"type"`
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	Attributes struct {
		CreatedAt    interface{} `json:"createdAt"`
		UpdatedAt    interface{} `json:"updatedAt"`
		ExternalSite string      `json:"externalSite"`
		ExternalID   string      `json:"externalId"`
	} `json:"attributes"`
	Relationships struct {
		Item struct {
			// Links struct {
			// 	Self    string `json:"self"`
			// 	Related string `json:"related"`
			// } `json:"links"`
			Data struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"item"`
	} `json:"relationships"`
}
