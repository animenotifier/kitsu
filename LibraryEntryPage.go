package kitsu

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
	response, requestError := Get(query)

	if requestError != nil {
		return nil, requestError
	}

	page := new(LibraryEntryPage)
	decodeError := response.Unmarshal(page)

	return page, decodeError
}
