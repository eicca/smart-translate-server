package data

import (
	"fmt"
	"log"
	"strings"

	"github.com/eicca/translate-server/cache"
)

// SuggestionReq contains:
// - Query: the text user types in auto-complete field
// - Locales: list of locales user is using
// - FallbackLocale: locale which will be returned if it's not possible
//   detect current locale in the range of `Locales`
type SuggestionReq struct {
	Locales        []Locale
	FallbackLocale Locale
	Query          string
}

// SuggestionResp returns detected locale and suggestions for the locale.
// NOTE: current chrome extensions is not using this.
type SuggestionResp struct {
	Source      Locale
	Suggestions []Suggestion
}

// Suggestion contains a text and a locale of this text.
type Suggestion struct {
	Text   string `json:"phrase"`
	Locale Locale `json:"locale"`
}

// NormalizeLocale check if the `sourceLocale` is among user `Locales`.
// If `Locales` don't contain the source locale, it will return `FallbackLocale`
func (s SuggestionReq) NormalizeLocale(sourceLocale Locale) Locale {
	for _, locale := range s.Locales {
		if locale == sourceLocale {
			return sourceLocale
		}
	}

	return s.FallbackLocale
}

// TargetLocale returns the next locale after the sourceLocale.
// TODO: it's stupid and suggestions should have different multiple locales.
func (s SuggestionReq) TargetLocale(sourceLocale Locale) Locale {
	for _, locale := range s.Locales {
		if locale != sourceLocale {
			return locale
		}
	}
	log.Println("ERROR: TargetLocale should always return a locale")
	return s.Locales[0]
}

// FromCache retrieves the response from cache.
// If it's not here the error will be returned.
func (s SuggestionReq) FromCache() (*[]Suggestion, error) {
	var resp []Suggestion
	err := cache.DefaultClient.Get(s, &resp)
	return &resp, err
}

// SaveCache saves the response to the cache.
func (s SuggestionReq) SaveCache(resp *[]Suggestion) error {
	return cache.DefaultClient.Set(s, resp)
}

// CacheKey defines which components are responsible for uniqueness
// of the response.
func (s SuggestionReq) CacheKey() string {
	locales := strings.Join(LocalesToStrings(s.Locales), "")
	return fmt.Sprintf("suggestions:%s:%s", locales, s.Query)
}
