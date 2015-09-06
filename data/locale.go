package data

// Locale is a two letters ISO locale.
type Locale string

// StringsAsLocales converts alice of strings into slices of locales.
func StringsAsLocales(in []string) (locales []Locale) {
	for _, s := range in {
		locales = append(locales, Locale(s))
	}
	return
}
