package kitsu

import "encoding/json"

// Auto-generated JSON definition from:
// https://mholt.github.io/json-to-go/

// AnimePage represents one page containing up to 20 anime objects.
type AnimePage struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Slug                string `json:"slug"`
			Synopsis            string `json:"synopsis"`
			CoverImageTopOffset int    `json:"coverImageTopOffset"`
			Titles              struct {
				En   string `json:"en"`
				EnJp string `json:"en_jp"`
				JaJp string `json:"ja_jp"`
			} `json:"titles"`
			CanonicalTitle    string   `json:"canonicalTitle"`
			AbbreviatedTitles []string `json:"abbreviatedTitles"`
			AverageRating     string   `json:"averageRating"`
			RatingFrequencies struct {
				Num2  string `json:"2"`
				Num3  string `json:"3"`
				Num4  string `json:"4"`
				Num5  string `json:"5"`
				Num6  string `json:"6"`
				Num7  string `json:"7"`
				Num8  string `json:"8"`
				Num9  string `json:"9"`
				Num10 string `json:"10"`
				Num11 string `json:"11"`
				Num12 string `json:"12"`
				Num13 string `json:"13"`
				Num14 string `json:"14"`
				Num15 string `json:"15"`
				Num16 string `json:"16"`
				Num17 string `json:"17"`
				Num18 string `json:"18"`
				Num19 string `json:"19"`
				Num20 string `json:"20"`
			} `json:"ratingFrequencies"`
			UserCount      int    `json:"userCount"`
			FavoritesCount int    `json:"favoritesCount"`
			StartDate      string `json:"startDate"`
			EndDate        string `json:"endDate"`
			PopularityRank int    `json:"popularityRank"`
			RatingRank     int    `json:"ratingRank"`
			AgeRating      string `json:"ageRating"`
			AgeRatingGuide string `json:"ageRatingGuide"`
			Subtype        string `json:"subtype"`
			PosterImage    struct {
				Tiny     string `json:"tiny"`
				Small    string `json:"small"`
				Medium   string `json:"medium"`
				Large    string `json:"large"`
				Original string `json:"original"`
			} `json:"posterImage"`
			CoverImage struct {
				Tiny     string `json:"tiny"`
				Small    string `json:"small"`
				Large    string `json:"large"`
				Original string `json:"original"`
			} `json:"coverImage"`
			EpisodeCount   int    `json:"episodeCount"`
			EpisodeLength  int    `json:"episodeLength"`
			YoutubeVideoID string `json:"youtubeVideoId"`
			ShowType       string `json:"showType"`
			Nsfw           bool   `json:"nsfw"`
		} `json:"attributes"`
		Relationships struct {
			Genres struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"genres"`
			Categories struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"categories"`
			Castings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"castings"`
			Installments struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"installments"`
			Mappings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mappings"`
			Reviews struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reviews"`
			MediaRelationships struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mediaRelationships"`
			Episodes struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"episodes"`
			StreamingLinks struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"streamingLinks"`
			AnimeProductions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeProductions"`
			AnimeCharacters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeCharacters"`
			AnimeStaff struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"animeStaff"`
		} `json:"relationships"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetAnimePage expects the usual query parameter and returns an AnimePage object instead of a raw string.
func GetAnimePage(query string) (*AnimePage, error) {
	body, requestError := Get(query)

	if requestError != nil {
		return nil, requestError[0]
	}

	page := new(AnimePage)
	decodeError := json.Unmarshal(body, page)

	return page, decodeError
}
