package kitsu

import (
	"encoding/json"
)

// Auto-generated JSON definition from:
// https://mholt.github.io/json-to-go/

// AnimePage represents one page containing up to 20 anime objects.
type AnimePage struct {
	Data     []*Anime `json:"data"`
	Included []struct {
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
	} `json:"included"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetAnimePage expects the usual query parameter and returns an AnimePage object instead of a raw string.
func GetAnimePage(query string) (*AnimePage, error) {
	body, requestError := Get(query)

	if requestError != nil {
		return nil, requestError[0]
	}

	page := new(AnimePage)
	decodeError := json.Unmarshal(body, page)

	return page, decodeError
}
