package glosbe

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/eicca/translate-server/data"
	"github.com/eicca/translate-server/gtranslate"
	"github.com/eicca/translate-server/httputils"
)

const (
	maxSuggestions = 5
	suggestURL     = "http://glosbe.com/ajax/phrasesAutosuggest"
)

type suggestResp []string

// Suggest returns a slice of suggestions for a request.
// glosbe returns suggestions only for the source locale.
func Suggest(req data.SuggestionReq) (*[]data.Suggestion, error) {
	if resp, err := req.FromCache(); err == nil {
		return resp, nil
	}

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

	rawData, err := httputils.Get(query)
	if err != nil {
		return nil, err
	}

	resp, err := parseSuggestResp(rawData, source)
	if err != nil {
		if cacheErr := req.SaveCache(resp); cacheErr != nil {
			log.Println(cacheErr)
		}
	}

	return resp, err
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

func parseSuggestResp(rawData []byte, locale data.Locale) (*[]data.Suggestion, error) {
	suggestResp := []string{}
	if err := json.Unmarshal(rawData, &suggestResp); err != nil {
		return nil, err
	}

	if len(suggestResp) > maxSuggestions {
		suggestResp = suggestResp[:maxSuggestions]
	}

	suggestions := []data.Suggestion{}
	for _, text := range suggestResp {
		s := data.Suggestion{Text: text, Locale: locale}
		suggestions = append(suggestions, s)
	}

	return &suggestions, nil
}
