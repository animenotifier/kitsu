package kitsu

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestGetAnime(t *testing.T) {
	anime, err := GetAnime("12268")

	assert.Nil(t, err)
	assert.NotNil(t, anime)
	assert.NotEqual(t, anime.ID, "")
	assert.NotEqual(t, anime.Attributes.Titles.EnJp, "")
	assert.NotEqual(t, anime.Link(), "")
	assert.NotEqual(t, anime.Attributes.PosterImage.Original, "")
	assert.NotEqual(t, FixImageURL(anime.Attributes.PosterImage.Original), "")
}

func TestGetAnimeCharactersForAnime(t *testing.T) {
	response, err := GetAnimeCharactersForAnime("12268")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Data)
	assert.NotEqual(t, response.Included, "")
}

func TestAnimePage(t *testing.T) {
	page, err := GetAnimePage("anime?page[limit]=5&page[offset]=0")

	assert.Nil(t, err)

	for _, anime := range page.Data {
		assert.NotEqual(t, anime.ID, "")
		assert.NotEqual(t, anime.Attributes.Titles.EnJp, "")
		assert.NotEqual(t, anime.Link(), "")
		assert.NotEqual(t, anime.Attributes.PosterImage.Original, "")
		assert.NotEqual(t, FixImageURL(anime.Attributes.PosterImage.Original), "")
	}
}

func TestMappingPage(t *testing.T) {
	page, err := GetMappingsPage("mappings?page[limit]=5&page[offset]=0&include=item")

	assert.Nil(t, err)

	for _, mapping := range page.Data {
		assert.NotEqual(t, mapping.ID, "")
		assert.NotEqual(t, mapping.Attributes.ExternalID, "")
		assert.NotEqual(t, mapping.Attributes.ExternalSite, "")
		assert.NotEqual(t, mapping.Relationships.Item.Data.ID, "")
	}
}

func TestStreamAnime(t *testing.T) {
	allAnime := StreamAnime()

	count := 0
	for anime := range allAnime {
		assert.NotEqual(t, anime.ID, "")
		assert.NotEqual(t, anime.Attributes.Titles.EnJp, "")

		count++

		if count >= 40 {
			break
		}
	}
}

func TestStreamMediaRelations(t *testing.T) {
	allRelations := StreamMediaRelations()

	count := 0
	for relation := range allRelations {
		assert.NotEqual(t, relation.ID, "")
		assert.NotEqual(t, relation.Attributes.Role, "")
		assert.NotEqual(t, relation.Relationships.Source.Data.ID, "")
		assert.NotEqual(t, relation.Relationships.Source.Data.Type, "")
		assert.NotEqual(t, relation.Relationships.Destination.Data.ID, "")
		assert.NotEqual(t, relation.Relationships.Destination.Data.Type, "")

		count++

		if count >= 40 {
			break
		}
	}
}

func TestCharacterPage(t *testing.T) {
	page, err := GetCharacterPage("characters?page[limit]=5&page[offset]=0")

	assert.Nil(t, err)

	for _, character := range page.Data {
		assert.NotEqual(t, character.ID, "")
		assert.NotEqual(t, character.Attributes.CanonicalName, "")
	}
}

func TestStreamCharacters(t *testing.T) {
	allCharacters := StreamCharacters()

	count := 0
	for character := range allCharacters {
		assert.NotEqual(t, character.ID, "")
		assert.NotEqual(t, character.Attributes.CanonicalName, "")

		count++

		if count >= 40 {
			break
		}
	}
}

func TestStreamMappings(t *testing.T) {
	allMappings := StreamMappings()

	count := 0
	for mapping := range allMappings {
		assert.NotEqual(t, mapping.ID, "")
		assert.NotEqual(t, mapping.Attributes.ExternalID, "")
		assert.NotEqual(t, mapping.Attributes.ExternalSite, "")
		assert.NotEqual(t, mapping.Relationships.Item.Data.ID, "")

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
	assert.Nil(t, err)
	assert.NotEqual(t, user.ID, "")
	assert.NotEqual(t, user.Attributes.Name, "")
}

func TestGetUserError(t *testing.T) {
	userName := "😀🙄😇"
	user, err := GetUser(userName)

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestStreamUserLibrary(t *testing.T) {
	userName := "Akyoto"
	user, _ := GetUser(userName)
	library := user.StreamLibraryEntries()

	count := 0
	for entry := range library {
		assert.NotNil(t, entry)
		assert.NotEqual(t, entry.Attributes.Status, "")

		assert.NotEqual(t, entry.ID, "")
		assert.NotEqual(t, entry.Attributes.Status, "")

		if entry.Anime != nil {
			assert.NotEqual(t, entry.Anime.Attributes.CanonicalTitle, "")
		}

		count++

		if count >= 40 {
			break
		}
	}
}
