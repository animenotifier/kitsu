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
	Links struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetAnimePage expects the usual query parameter and returns an AnimePage object instead of a raw string.
func GetAnimePage(query string) (*AnimePage, error) {
	response, requestError := Get(query)

	if requestError != nil {
		return nil, requestError
	}

	page := new(AnimePage)
	decodeError := response.Unmarshal(page)

	return page, decodeError
}
