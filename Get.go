package kitsu

import "github.com/parnurzeal/gorequest"

const (
	urlBase    = "https://kitsu.io/api/edge/"
	acceptType = "application/vnd.api+json"
)

// Get performs a GET request to the Kitsu API.
func Get(query string) ([]byte, []error) {
	_, body, err := gorequest.New().Get(urlBase+query).Set("Accept", acceptType).EndBytes()
	return body, err
}
