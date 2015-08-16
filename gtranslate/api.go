package gtranslate

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	translateURL = "https://www.googleapis.com/language/translate/v2"
	webURL       = "https://translate.google.com/#"
)

func get(u *url.URL) ([]byte, error) {
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
