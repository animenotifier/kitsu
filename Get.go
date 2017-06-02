package kitsu

import "github.com/parnurzeal/gorequest"

const (
	// APIBaseURL is the prefix for all API requests.
	APIBaseURL = "https://kitsu.io/api/edge/"

	acceptType = "application/vnd.api+json"
)

// Get performs a GET request to the Kitsu API.
func Get(query string) ([]byte, []error) {
	_, body, err := gorequest.New().Get(APIBaseURL+query).Set("Accept", acceptType).EndBytes()
	return body, err
}
