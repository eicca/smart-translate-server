package glosbe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eicca/translate-server/data"
	"github.com/eicca/translate-server/gtranslate"
)

const (
	suggestURL = "http://glosbe.com/ajax/phrasesAutosuggest"
)

type suggestResp []string

// Suggest returns a slice of suggestions for a request.
// glosbe returns suggestions only for the source locale.
func Suggest(req data.SuggestionReq) ([]data.Suggestion, error) {
	source, err := gtranslate.Detect(req)
	if err != nil {
		return nil, err
	}

	source = req.NormalizeLocale(source)
	target := req.TargetLocale(source)

	query, err := makeSuggestQuery(req.Query, source, target)
	if err != nil {
		return nil, err
	}

	rawData, err := get(query)
	if err != nil {
		return nil, err
	}

	return parseSuggestResp(rawData, source)
}

func makeSuggestQuery(query string, source data.Locale, target data.Locale) (*url.URL, error) {
	u, err := url.Parse(suggestURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("from", source.String())
	q.Set("dest", target.String())
	q.Set("phrase", query)
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

// TODO duplication with gtranslate.get
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
