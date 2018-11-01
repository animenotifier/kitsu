package kitsu

import (
	"errors"
)

// GetUser returns the user with the given nickname.
func GetUser(userName string) (*User, error) {
	response, requestError := Get("users?filter[slug]=" + userName)

	if requestError != nil {
		return nil, requestError
	}

	user := new(UserResponse)
	decodeError := response.Unmarshal(user)

	if len(user.Data) == 0 {
		return nil, errors.New("User " + userName + " could not be found on Kitsu")
	}

	return user.Data[0], decodeError
}
