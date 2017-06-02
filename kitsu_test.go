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
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.En)
	}
}

func TestAllAnime(t *testing.T) {
	color.Yellow("Fetching all anime (stops after 60 anime)")

	allAnime := AllAnime()

	count := 0
	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.En)
		count++

		if count >= 60 {
			close(allAnime)
			break
		}
	}
}
