package kitsu

// GetAnimeCharactersForAnime ...
func GetAnimeCharactersForAnime(animeID string) (*AnimeCharactersResponse, error) {
	response, requestError := Get("anime/" + animeID + "?include=animeCharacters,animeCharacters.character")

	if requestError != nil {
		return nil, requestError
	}

	characters := new(AnimeCharactersResponse)
	response.Unmarshal(characters)

	return characters, nil
}
