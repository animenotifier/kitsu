package kitsu

import (
	"strings"
	"time"
)

// StreamCharacters returns a stream of all character objects (async).
// Be very careful to only use this function once as each
// call will start a new goroutine requesting the whole data.
func StreamCharacters() chan *Character {
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
