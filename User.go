package kitsu

import "time"

// User ...
type User struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
		Name               string      `json:"name"`
		PastNames          []string    `json:"pastNames"`
		About              string      `json:"about"`
		Bio                string      `json:"bio"`
		AboutFormatted     interface{} `json:"aboutFormatted"`
		Location           string      `json:"location"`
		WaifuOrHusbando    interface{} `json:"waifuOrHusbando"`
		FollowersCount     int         `json:"followersCount"`
		FacebookID         string      `json:"facebookId"`
		FollowingCount     int         `json:"followingCount"`
		LifeSpentOnAnime   int         `json:"lifeSpentOnAnime"`
		Birthday           interface{} `json:"birthday"`
		Gender             string      `json:"gender"`
		CommentsCount      int         `json:"commentsCount"`
		FavoritesCount     int         `json:"favoritesCount"`
		LikesGivenCount    int         `json:"likesGivenCount"`
		ReviewsCount       int         `json:"reviewsCount"`
		LikesReceivedCount int         `json:"likesReceivedCount"`
		PostsCount         int         `json:"postsCount"`
		RatingsCount       int         `json:"ratingsCount"`
		ProExpiresAt       interface{} `json:"proExpiresAt"`
		Title              string      `json:"title"`
		ProfileCompleted   bool        `json:"profileCompleted"`
		FeedCompleted      bool        `json:"feedCompleted"`
		Website            string      `json:"website"`
		Avatar             struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"avatar"`
		CoverImage struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"coverImage"`
		RatingSystem string `json:"ratingSystem"`
		Theme        string `json:"theme"`
	} `json:"attributes"`
	Relationships struct {
		Waifu struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"waifu"`
		PinnedPost struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"pinnedPost"`
		Followers struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"followers"`
		Following struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"following"`
		Blocks struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"blocks"`
		LinkedAccounts struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"linkedAccounts"`
		ProfileLinks struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"profileLinks"`
		MediaFollows struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"mediaFollows"`
		UserRoles struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"userRoles"`
		LibraryEntries struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"libraryEntries"`
		Favorites struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"favorites"`
		Reviews struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"reviews"`
		Stats struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"stats"`
		NotificationSettings struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"notificationSettings"`
		OneSignalPlayers struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"oneSignalPlayers"`
	} `json:"relationships"`
}
