package kitsu

import (
	"strings"
)

// FixImageURL removes the right-most part, e.g. "?1416336000" from image URLs.
// Additionally it also remove the protocol ("https:")
func FixImageURL(url string) string {
	url = strings.TrimPrefix(url, "https:")
	questionMark := strings.IndexByte(url, '?')

	if questionMark == -1 {
		return url
	}

	return url[:questionMark]
}
