package kitsu

// CharacterPage represents one page containing up to 10 character objects.
type CharacterPage struct {
	Data []*Character `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetCharacterPage expects the usual query parameter and returns an CharacterPage object instead of a raw string.
func GetCharacterPage(query string) (*CharacterPage, error) {
	response, requestError := Get(query)

	if requestError != nil {
		return nil, requestError
	}

	page := new(CharacterPage)
	decodeError := response.Unmarshal(page)

	return page, decodeError
}
