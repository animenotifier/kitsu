package kitsu

import "encoding/json"

// GetAnimeCharactersForAnime ...
func GetAnimeCharactersForAnime(animeID string) (*AnimeCharactersResponse, error) {
	body, requestError := Get("anime/" + animeID + "?include=animeCharacters,animeCharacters.character")

	if requestError != nil {
		return nil, requestError[0]
	}

	resp := &AnimeCharactersResponse{}
	json.Unmarshal(body, &resp)

	return resp, nil
}
