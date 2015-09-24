package httputils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get sends a get request to provided `url` and reads a response.
func Get(u *url.URL) ([]byte, error) {
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
