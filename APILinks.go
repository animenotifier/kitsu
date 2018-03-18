package kitsu

// APILinks ...
type APILinks struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}
