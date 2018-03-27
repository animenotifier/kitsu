package kitsu

// Auto-generated JSON definition from:
// https://mholt.github.io/json-to-go/

// Character is a character in the Kitsu API.
// Example: https://kitsu.io/api/edge/characters/1
type Character struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	Attributes struct {
		Slug  string `json:"slug"`
		Names struct {
			En   string `json:"en"`
			JaJp string `json:"ja_jp"`
		} `json:"names"`
		CanonicalName string   `json:"canonicalName"`
		OtherNames    []string `json:"otherNames"`
		MalID         int      `json:"malId"`
		Description   string   `json:"description"`
		Image         struct {
			Original string `json:"original"`
		} `json:"image"`
	} `json:"attributes"`
	// Relationships struct {
	// 	PrimaryMedia struct {
	// 		Links struct {
	// 			Self    string `json:"self"`
	// 			Related string `json:"related"`
	// 		} `json:"links"`
	// 	} `json:"primaryMedia"`
	// 	Castings struct {
	// 		Links struct {
	// 			Self    string `json:"self"`
	// 			Related string `json:"related"`
	// 		} `json:"links"`
	// 	} `json:"castings"`
	// } `json:"relationships"`
}
