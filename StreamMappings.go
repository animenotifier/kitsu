package kitsu

import (
	"strings"
	"time"
)

// StreamMappings returns a stream of all mappings (async).
func StreamMappings() chan *Mapping {
	channel := make(chan *Mapping)
	url := "mappings?page[limit]=20&page[offset]=0&include=item"
	ticker := time.NewTicker(500 * time.Millisecond)
	rateLimit := ticker.C

	go func() {
		defer close(channel)
		defer ticker.Stop()

		for {
			page, err := GetMappingsPage(url)

			if err != nil {
				panic(err)
			}

			// Feed mapping data from current page to the stream
			for _, mapping := range page.Data {
				channel <- mapping
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
