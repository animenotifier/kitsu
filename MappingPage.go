package kitsu

// MappingPage ...
// Example: https://kitsu.io/api/edge/mappings/?include=item
type MappingPage struct {
	Data  []*Mapping `json:"data"`
	Links APILinks   `json:"links"`
}

// GetMappingPage expects the usual query parameter and returns an MappingPage object.
func GetMappingPage(query string) (*MappingPage, error) {
	response, err := Get(query)

	if err != nil {
		return nil, err
	}

	page := &MappingPage{}
	err = response.Unmarshal(page)

	return page, err
}
