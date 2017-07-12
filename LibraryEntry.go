package kitsu

import "time"

// Relationship ...
type Relationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// LibraryEntry ...
type LibraryEntry struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt      time.Time   `json:"createdAt"`
		UpdatedAt      time.Time   `json:"updatedAt"`
		Status         string      `json:"status"`
		Progress       int         `json:"progress"`
		VolumesOwned   int         `json:"volumesOwned"`
		Reconsuming    bool        `json:"reconsuming"`
		ReconsumeCount int         `json:"reconsumeCount"`
		Notes          interface{} `json:"notes"`
		Private        bool        `json:"private"`
		ProgressedAt   interface{} `json:"progressedAt"`
		StartedAt      interface{} `json:"startedAt"`
		FinishedAt     interface{} `json:"finishedAt"`
		Rating         string      `json:"rating"`
		RatingTwenty   interface{} `json:"ratingTwenty"`
	} `json:"attributes"`
	Relationships struct {
		User struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"user"`
		Anime struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
			Data *Relationship `json:"data"`
		} `json:"anime"`
		Manga struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"manga"`
		Drama struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"drama"`
		Review struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"review"`
		MediaReaction struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"mediaReaction"`
		Media struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"media"`
		Unit struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"unit"`
		NextUnit struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"nextUnit"`
	} `json:"relationships"`

	// Custom fields, added by this client.
	// Doesn't really exist in the API
	Anime *Anime
}
