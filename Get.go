package kitsu

import "github.com/aerogo/http/client"

const (
	// APIBaseURL is the prefix for all API requests.
	APIBaseURL = "https://kitsu.io/api/edge/"

	acceptType = "application/vnd.api+json"
)

// Get performs a GET request to the Kitsu API.
func Get(query string) (*client.Response, error) {
	response, err := client.Get(APIBaseURL+query).Header("Accept", acceptType).End()
	return response, err
}
