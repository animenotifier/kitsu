package kitsu

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
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

func TestAllAnime(t *testing.T) {
	color.Yellow("Fetching all anime (stops after 60 anime)")

	allAnime := StreamAnime()

	count := 0
	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.EnJp)
		count++

		if count >= 40 {
			// close(allAnime)
			break
		}
	}
}

func TestAllAnimeMappings(t *testing.T) {
	color.Yellow("Fetching all anime with mappings (stops after 60 anime)")

	allAnime := StreamAnimeWithMappings()

	count := 0
	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]")

		for _, mapping := range anime.Mappings {
			fmt.Println(mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)
		}

		count++

		if count >= 40 {
			// close(allAnime)
			break
		}
	}
}

func TestAllCharacters(t *testing.T) {
	color.Yellow("Fetching all characters (stops after 60 characters)")

	allCharacters := StreamCharacters()

	count := 0
	for character := range allCharacters {
		fmt.Println("[", character.ID, "]", character.Attributes.Name)
		count++

		if count >= 60 {
			close(allCharacters)
			break
		}
	}
}
