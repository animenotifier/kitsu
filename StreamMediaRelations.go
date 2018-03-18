package kitsu

import (
	"strings"
	"time"
)

// StreamMediaRelations returns a stream of all anime relations (async).
func StreamMediaRelations() chan *MediaRelation {
	channel := make(chan *MediaRelation)
	url := "media-relationships?page[limit]=20&include=source,destination"
	ticker := time.NewTicker(1000 * time.Millisecond)
	rateLimit := ticker.C

	go func() {
		defer close(channel)
		defer ticker.Stop()

		for {
			page, err := GetMediaRelations(url)

			if err != nil {
				panic(err)
			}

			// Feed media relation data from current page to the stream
			for _, relation := range page.Data {
				channel <- relation
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
