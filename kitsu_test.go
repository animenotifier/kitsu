package kitsu

import (
	"fmt"
	"testing"
)

func TestAnimePage(t *testing.T) {
	fmt.Println("Fetching first 5 anime titles")

	page, err := GetAnimePage("anime?page[limit]=5&page[offset]=0")

	if err != nil {
		panic(err)
	}

	for _, anime := range page.Data {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.En)
	}
}

func TestAllAnime(t *testing.T) {
	fmt.Println("Fetching all anime")

	allAnime := AllAnime()

	for anime := range allAnime {
		fmt.Println("[", anime.ID, "]", anime.Attributes.Titles.En)
	}
}
