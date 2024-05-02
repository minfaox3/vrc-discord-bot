package vrchatapi

func (v *Endpoint) URL() string {
	return v.url
}

func (v *Endpoint) Method() MethodType {
	return v.method
}

var (
	endpointLogin = Endpoint{
		url:    "https://api.vrchat.cloud/api/1/auth/user",
		method: GET,
	}

	endpointEmail2FA = Endpoint{
		url:    "https://api.vrchat.cloud/api/1/auth/twofactorauth/emailotp/verify",
		method: POST,
	}

	endpointFriendsOnline = Endpoint{
		url:    "https://api.vrchat.cloud/api/1/auth/user/friends?offline=false",
		method: GET,
	}

	endpointFriendsOffline = Endpoint{
		url:    "https://api.vrchat.cloud/api/1/auth/user/friends?offline=true",
		method: GET,
	}

	endpointLogout = Endpoint{
		url:    "https://api.vrchat.cloud/api/1/logout",
		method: PUT,
	}
)
