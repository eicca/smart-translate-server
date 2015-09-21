package glosbe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eicca/translate-server/data"
)

const (
	suggestURL = "http://glosbe.com/ajax/phrasesAutosuggest"
)

type suggestResp []string

func Suggest(req data.SuggestionReq) ([]data.Suggestion, error) {
	query, err := makeSuggestQuery(req)
	if err != nil {
		return nil, err
	}

	rawData, err := get(query)
	if err != nil {
		return nil, err
	}

	return parseSuggestResp(rawData, req.Source)
}

func makeSuggestQuery(req data.SuggestionReq) (*url.URL, error) {
	u, err := url.Parse(suggestURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("from", req.Source.String())
	q.Set("dest", req.Target.String())
	q.Set("phrase", req.Query)
	u.RawQuery = q.Encode()
	return u, nil
}

func parseSuggestResp(rawData []byte, locale data.Locale) ([]data.Suggestion, error) {
	suggestResp := []string{}
	if err := json.Unmarshal(rawData, &suggestResp); err != nil {
		return nil, err
	}

	suggestions := []data.Suggestion{}
	for _, text := range suggestResp {
		s := data.Suggestion{Text: text, Locale: locale}
		suggestions = append(suggestions, s)
	}

	return suggestions, nil
}

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
