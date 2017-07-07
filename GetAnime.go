package kitsu

import "encoding/json"

// GetAnime returns the anime for the given ID.
func GetAnime(id string) (*Anime, error) {
	body, requestError := Get("anime/" + id)

	if requestError != nil {
		return nil, requestError[0]
	}

	response := new(AnimeResponse)
	decodeError := json.Unmarshal(body, response)

	return response.Data, decodeError
}
