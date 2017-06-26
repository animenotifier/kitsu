package kitsu

import (
	"strings"
	"time"
)

// Auto-generated JSON definition from:
// https://mholt.github.io/json-to-go/

// Character is a character in the Kitsu API.
type Character struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		MalID       int    `json:"malId"`
		Description string `json:"description"`
		Image       struct {
			Original string `json:"original"`
		} `json:"image"`
	} `json:"attributes"`
	Relationships struct {
		PrimaryMedia struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"primaryMedia"`
		Castings struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"castings"`
	} `json:"relationships"`
}

// AllCharacters returns a stream of all character objects (async).
// Be very careful to only use this function once as each
// call will start a new goroutine requesting the whole data.
func AllCharacters() chan *Character {
	channel := make(chan *Character)
	url := "characters?page[limit]=20&page[offset]=0"
	ticker := time.NewTicker(500 * time.Millisecond)
	rateLimit := ticker.C

	go func() {
		defer close(channel)
		defer ticker.Stop()

		for {
			page, err := GetCharacterPage(url)

			if err != nil {
				panic(err)
			}

			// Feed character data from current page to the stream
			for _, character := range page.Data {
				channel <- character
			}

			nextURL := page.Links.Next

			// Did we reach the end?
			if nextURL == "" {
				break
			}

			// Cut off API base URL
			nextURL = strings.TrimPrefix(nextURL, APIBaseURL)

			// Continue with the next page
			url = nextURL

			// Wait for rate limiter to allow the next request
			<-rateLimit
		}
	}()

	return channel
}
