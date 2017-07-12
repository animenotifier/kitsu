package kitsu

import (
	"encoding/json"
)

// LibraryEntryPage ...
type LibraryEntryPage struct {
	Data     []*LibraryEntry `json:"data"`
	Included []*Anime        `json:"included"`
	Meta     struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetLibraryEntryPage ...
func GetLibraryEntryPage(query string) (*LibraryEntryPage, error) {
	body, requestError := Get(query)

	if requestError != nil {
		return nil, requestError[0]
	}

	page := new(LibraryEntryPage)
	decodeError := json.Unmarshal(body, page)

	return page, decodeError
}
