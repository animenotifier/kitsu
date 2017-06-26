package kitsu

import (
	"strings"
)

// FixImageURL removes the right-most part, e.g. "?1416336000" from image URLs.
func FixImageURL(url string) string {
	questionMark := strings.IndexByte(url, '?')

	if questionMark == -1 {
		return url
	}

	return url[:questionMark]
}
