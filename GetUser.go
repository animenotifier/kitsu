package kitsu

import (
	"encoding/json"
	"errors"
)

// GetUser returns the user with the given nickname.
func GetUser(userName string) (*User, error) {
	body, requestError := Get("users?filter[name]=" + userName)

	if requestError != nil {
		return nil, requestError[0]
	}

	response := new(UserResponse)
	decodeError := json.Unmarshal(body, response)

	if len(response.Data) == 0 {
		return nil, errors.New("User " + userName + " could not be found on Kitsu")
	}

	return response.Data[0], decodeError
}
