package data

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
	// TODO replace panic with logging.
	panic("TargetLocale should always return a locale")
}
