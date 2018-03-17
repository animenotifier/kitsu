package kitsu

// StreamAnimeWithMappings returns a stream of all anime objects with mappings.
// func StreamAnimeWithMappings() chan *Anime {
// 	channel := make(chan *Anime)
// 	url := "anime?page[limit]=20&page[offset]=0&include=mappings"
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	rateLimit := ticker.C

// 	go func() {
// 		defer close(channel)
// 		defer ticker.Stop()

// 		for {
// 			page, err := GetAnimePage(url)

// 			if err != nil {
// 				panic(err)
// 			}

// 			mappings := map[string]*Mapping{}

// 			for _, mapping := range page.Included {
// 				mappings[mapping.ID] = mapping
// 			}

// 			// Feed anime data from current page to the stream
// 			for _, anime := range page.Data {
// 				for _, mapping := range anime.Relationships.Mappings.Data {
// 					anime.Mappings = append(anime.Mappings, mappings[mapping.ID])
// 				}

// 				channel <- anime
// 			}

// 			nextURL := page.Links.Next

// 			// Did we reach the end?
// 			if nextURL == "" {
// 				break
// 			}

// 			// Cut off API base URL
// 			nextURL = strings.TrimPrefix(nextURL, APIBaseURL)

// 			// Continue with the next page
// 			url = nextURL

// 			// Wait for rate limiter to allow the next request
// 			<-rateLimit
// 		}
// 	}()

// 	return channel
// }
