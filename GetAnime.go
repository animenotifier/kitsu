package kitsu

// GetAnime returns the anime for the given ID.
func GetAnime(id string) (*Anime, error) {
	response, requestError := Get("anime/" + id)

	if requestError != nil {
		return nil, requestError
	}

	anime := new(AnimeResponse)
	decodeError := response.Unmarshal(anime)

	return anime.Data, decodeError
}
