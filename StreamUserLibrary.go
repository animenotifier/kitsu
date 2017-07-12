package kitsu

import (
	"strings"
	"time"
)

// StreamLibraryEntries returns a stream of all library entries (async).
func (user *User) StreamLibraryEntries() chan *LibraryEntry {
	channel := make(chan *LibraryEntry)
	url := "users/" + user.ID + "/library-entries?page[limit]=10&page[offset]=0&include=anime"
	ticker := time.NewTicker(500 * time.Millisecond)
	rateLimit := ticker.C

	go func() {
		defer close(channel)
		defer ticker.Stop()

		for {
			page, err := GetLibraryEntryPage(url)

			if err != nil {
				panic(err)
			}

			// Feed entry data from current page to the stream
			includedAnime := map[string]*Anime{}

			for _, anime := range page.Included {
				includedAnime[anime.ID] = anime
			}

			// Feed entry data from current page to the stream
			for _, entry := range page.Data {
				animeInfo := entry.Relationships.Anime.Data

				if animeInfo != nil {
					entry.Anime = includedAnime[animeInfo.ID]
				}

				channel <- entry
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
