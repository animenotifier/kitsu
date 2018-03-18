package kitsu

// MappingsPage ...
// Example: https://kitsu.io/api/edge/mappings/?include=item
type MappingsPage struct {
	Data  []*Mapping `json:"data"`
	Links APILinks   `json:"links"`
}

// GetMappingsPage expects the usual query parameter and returns an MappingPage object.
func GetMappingsPage(query string) (*MappingsPage, error) {
	response, err := Get(query)

	if err != nil {
		return nil, err
	}

	page := &MappingsPage{}
	err = response.Unmarshal(page)

	return page, err
}
