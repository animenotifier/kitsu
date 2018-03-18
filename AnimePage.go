package kitsu

// Auto-generated JSON definition from:
// https://mholt.github.io/json-to-go/

// AnimePage represents one page containing up to 20 anime objects.
type AnimePage struct {
	Data     []*Anime   `json:"data"`
	Included []*Mapping `json:"included"`
	Meta     struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links APILinks `json:"links"`
}

// GetAnimePage expects the usual query parameter and returns an AnimePage object.
func GetAnimePage(query string) (*AnimePage, error) {
	response, err := Get(query)

	if err != nil {
		return nil, err
	}

	page := &AnimePage{}
	err = response.Unmarshal(page)

	return page, err
}
