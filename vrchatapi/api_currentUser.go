package vrchatapi

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func CurrentUserData(username, password, userAgent string, cookie []*http.Cookie) (bool, CurrentUser, []*http.Cookie, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(username) + ":" + url.QueryEscape(password)))
	headers := map[string]string{
		"Authorization": "Basic " + auth,
	}
	status, data, cookie, err := request(endpointLogin, userAgent, nil, &headers, cookie)
	if err != nil {
		return false, CurrentUser{}, cookie, err
	}
	if status != 200 {
		return false, CurrentUser{}, cookie, errors.New(string(data))
	}

	var fadata Need2FA
	err = json.Unmarshal(data, &fadata)
	if err != nil {
		return false, CurrentUser{}, cookie, err
	}
	if len(fadata.RequiresTwoFactorAuth) == 0 {
		var currentUser CurrentUser
		err = json.Unmarshal(data, &currentUser)
		if err != nil {
			return false, CurrentUser{}, cookie, err
		}
		return true, currentUser, cookie, nil
	}
	return false, CurrentUser{}, cookie, nil
}
