package kitsu

// Mapping ...
type Mapping struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		ExternalSite string `json:"externalSite"`
		ExternalID   string `json:"externalId"`
	} `json:"attributes"`
	Relationships struct {
		Media struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"media"`
	} `json:"relationships"`
}
