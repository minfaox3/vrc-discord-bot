package vrchatapi

import (
	"io"
	"net/http"
)

func request(endpoint Endpoint, userAgent string, body io.Reader, header *map[string]string, cookie []*http.Cookie) (int, []byte, []*http.Cookie, error) {
	req, err := http.NewRequest(string(endpoint.method), endpoint.url, body)
	if err != nil {
		return 0, nil, nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	if header != nil {
		for k, v := range *header {
			req.Header.Set(k, v)
		}
	}
	for _, c := range cookie {
		req.AddCookie(c)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}

	defer resp.Body.Close()

	cookie = append(cookie, resp.Cookies()...)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}

	return resp.StatusCode, data, cookie, nil
}
