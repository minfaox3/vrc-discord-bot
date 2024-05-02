package vrchatapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func Login(username, password, userAgent string) (LoginResult, Need2FA, CurrentUser, []*http.Cookie, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(username) + ":" + url.QueryEscape(password)))
	headers := map[string]string{
		"Authorization": "Basic " + auth,
	}
	status, data, cookie, err := request(endpointLogin, userAgent, nil, &headers, nil)
	if err != nil {
		return Failed, Need2FA{}, CurrentUser{}, cookie, err
	}
	if status != 200 {
		return Failed, Need2FA{}, CurrentUser{}, cookie, errors.New(string(data))
	}

	var fadata Need2FA
	err = json.Unmarshal(data, &fadata)
	if err != nil {
		return Failed, Need2FA{}, CurrentUser{}, cookie, err
	}
	if len(fadata.RequiresTwoFactorAuth) == 0 {
		var currentUser CurrentUser
		err = json.Unmarshal(data, &currentUser)
		if err != nil {
			return Failed, Need2FA{}, CurrentUser{}, cookie, err
		}
		return Succeeded, Need2FA{}, currentUser, cookie, nil
	}
	return Wait2FA, fadata, CurrentUser{}, cookie, nil
}

func Mail2FA(code string, userAgent string, cookie []*http.Cookie) (bool, []*http.Cookie, error) {
	body, err := json.Marshal(Code2FA{Code: code})
	if err != nil {
		return false, cookie, err
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	status, data, cookie, err := request(endpointEmail2FA, userAgent, bytes.NewBuffer(body), &headers, cookie)
	if err != nil {
		return false, cookie, err
	}
	if status != 200 {
		return false, cookie, errors.New(strconv.Itoa(status) + ": " + string(data))
	}
	return true, cookie, nil
}

func Logout(userAgent string, cookie []*http.Cookie) (bool, []*http.Cookie, error) {
	status, data, cookie, err := request(endpointLogout, userAgent, nil, nil, cookie)
	if err != nil {
		return false, nil, err
	}
	if status != 200 {
		return false, cookie, errors.New(strconv.Itoa(status) + ": " + string(data))
	}
	return true, cookie, nil
}
