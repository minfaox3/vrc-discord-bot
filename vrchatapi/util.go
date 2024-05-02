package vrchatapi

func CheckUserAgent(userAgent string) string {
	if userAgent == "" {
		return "application/1.00 user@example.com"
	} else {
		return userAgent
	}
}
