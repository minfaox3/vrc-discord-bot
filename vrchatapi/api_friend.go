package vrchatapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func FriendOnline(userAgent string, cookie []*http.Cookie) ([]FriendUser, error) {
	status, data, cookie, err := request(endpointFriendsOnline, userAgent, nil, nil, cookie)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New(strconv.Itoa(status) + ":" + string(data))
	}
	var friends []FriendUser
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func FriendOffline(userAgent string, cookie []*http.Cookie) ([]FriendUser, error) {
	status, data, cookie, err := request(endpointFriendsOffline, userAgent, nil, nil, cookie)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New(strconv.Itoa(status) + ":" + string(data))
	}
	var friends []FriendUser
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
