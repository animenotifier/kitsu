package kitsu

// GetAnimeCharactersForAnime ...
func GetAnimeCharactersForAnime(animeID string) (*AnimeCharactersResponse, error) {
	response, requestError := Get("anime/" + animeID + "?include=animeCharacters,animeCharacters.character")

	if requestError != nil {
		return nil, requestError
	}

	characters := &AnimeCharactersResponse{}
	err := response.Unmarshal(characters)

	if err != nil {
		return nil, err
	}

	return characters, nil
}
