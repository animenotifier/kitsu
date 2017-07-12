package kitsu

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestAnimePage(t *testing.T) {
	color.Yellow("Fetching first 5 anime titles")

	page, err := GetAnimePage("anime?page[limit]=5&page[offset]=0")

	if err != nil {
		panic(err)
	}

	for _, anime := range page.Data {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.EnJp)
	}
}

func TestCharacterPage(t *testing.T) {
	color.Yellow("Fetching first 5 anime titles")

	page, err := GetCharacterPage("characters?page[limit]=5&page[offset]=0")

	if err != nil {
		panic(err)
	}

	for _, character := range page.Data {
		fmt.Println("[", character.ID, "]", character.Attributes.Name)
	}
}

func TestStreamAnime(t *testing.T) {
	color.Yellow("Fetching all anime (stops after 40 anime)")

	allAnime := StreamAnime()

	count := 0
	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.EnJp)
		count++

		if count >= 40 {
			break
		}
	}
}

func TestStreamAnimeMappings(t *testing.T) {
	color.Yellow("Fetching all anime with mappings (stops after 40 anime)")

	allAnime := StreamAnimeWithMappings()

	count := 0
	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]")

		for _, mapping := range anime.Mappings {
			fmt.Println(mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
		}

		count++

		if count >= 40 {
			break
		}
	}
}

func TestStreamCharacters(t *testing.T) {
	color.Yellow("Fetching all characters (stops after 40 characters)")

	allCharacters := StreamCharacters()

	count := 0
	for character := range allCharacters {
		fmt.Println("[", character.ID, "]", character.Attributes.Name)
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
	color.Yellow("Fetching all library entries (stops after 40 entries)")

	userName := "Akyoto"
	user, _ := GetUser(userName)
	library := user.StreamLibraryEntries()

	count := 0
	for entry := range library {
		assert.NotNil(t, entry)
		assert.NotEmpty(t, entry.Attributes.Status)

		if entry.Anime != nil {
			fmt.Println("[", entry.ID, "]", entry.Attributes.Status, "|", entry.Anime.Attributes.CanonicalTitle)
		} else {
			fmt.Println("[", entry.ID, "]", entry.Attributes.Status)
		}

		count++

		if count >= 40 {
			break
		}
	}
}
