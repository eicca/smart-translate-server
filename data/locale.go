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

func LocalesToStrings(in []Locale) (out []string) {
	for _, l := range in {
		out = append(out, l.String())
	}
	return
}

func (l Locale) String() string {
	return string(l)
}
