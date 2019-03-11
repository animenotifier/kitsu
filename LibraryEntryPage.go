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
	response, err := Get(query)

	if err != nil {
		return nil, err
	}

	page := &LibraryEntryPage{}
	err = response.Unmarshal(page)

	return page, err
}
