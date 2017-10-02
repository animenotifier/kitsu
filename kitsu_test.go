package kitsu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnime(t *testing.T) {
	anime, err := GetAnime("12268")

	assert.NoError(t, err)
	assert.NotNil(t, anime)
	assert.NotEmpty(t, anime.ID)
	assert.NotEmpty(t, anime.Attributes.Titles.EnJp)
	assert.NotEmpty(t, anime.Link())
	assert.NotEmpty(t, anime.Attributes.PosterImage.Original)
	assert.NotEmpty(t, FixImageURL(anime.Attributes.PosterImage.Original))
}

func TestGetAnimeCharactersForAnime(t *testing.T) {
	response, err := GetAnimeCharactersForAnime("12268")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Data)
	assert.NotEmpty(t, response.Included)
}

func TestAnimePage(t *testing.T) {
	page, err := GetAnimePage("anime?page[limit]=5&page[offset]=0")

	assert.NoError(t, err)

	for _, anime := range page.Data {
		assert.NotEmpty(t, anime.ID)
		assert.NotEmpty(t, anime.Attributes.Titles.EnJp)
		assert.NotEmpty(t, anime.Link())
		assert.NotEmpty(t, anime.Attributes.PosterImage.Original)
		assert.NotEmpty(t, FixImageURL(anime.Attributes.PosterImage.Original))
	}
}

func TestStreamAnime(t *testing.T) {
	allAnime := StreamAnime()

	count := 0
	for anime := range allAnime {
		assert.NotEmpty(t, anime.ID)
		assert.NotEmpty(t, anime.Attributes.Titles.EnJp)

		count++

		if count >= 40 {
			break
		}
	}
}

func TestStreamAnimeMappings(t *testing.T) {
	allAnime := StreamAnimeWithMappings()

	count := 0
	for anime := range allAnime {
		assert.NotEmpty(t, anime.ID)

		for _, mapping := range anime.Mappings {
			assert.NotEmpty(t, mapping.Attributes.ExternalSite)
			assert.NotEmpty(t, mapping.Attributes.ExternalID)
		}

		count++

		if count >= 40 {
			break
		}
	}
}

func TestCharacterPage(t *testing.T) {
	page, err := GetCharacterPage("characters?page[limit]=5&page[offset]=0")

	assert.NoError(t, err)

	for _, character := range page.Data {
		assert.NotEmpty(t, character.ID)
		assert.NotEmpty(t, character.Attributes.Name)
	}
}

func TestStreamCharacters(t *testing.T) {
	allCharacters := StreamCharacters()

	count := 0
	for character := range allCharacters {
		assert.NotEmpty(t, character.ID)
		assert.NotEmpty(t, character.Attributes.Name)

		count++

		if count >= 40 {
			break
		}
	}
}

func TestGetUser(t *testing.T) {
	userName := "Akyoto"
	user, err := GetUser(userName)

	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Attributes.Name)
}

func TestGetUserError(t *testing.T) {
	userName := "404"
	user, err := GetUser(userName)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestStreamUserLibrary(t *testing.T) {
	userName := "Akyoto"
	user, _ := GetUser(userName)
	library := user.StreamLibraryEntries()

	count := 0
	for entry := range library {
		assert.NotNil(t, entry)
		assert.NotEmpty(t, entry.Attributes.Status)

		assert.NotEmpty(t, entry.ID)
		assert.NotEmpty(t, entry.Attributes.Status)

		if entry.Anime != nil {
			assert.NotEmpty(t, entry.Anime.Attributes.CanonicalTitle)
		}

		count++

		if count >= 40 {
			break
		}
	}
}
