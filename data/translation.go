package data

// MultiTranslationReq contains translation request for set of locales.
type MultiTranslationReq struct {
	Source  Locale
	Targets []Locale
	Query   string
}

// TranslationReq contains translation request for one locale.
type TranslationReq struct {
	Source Locale
	Target Locale
	Query  string
}

// MultiTranslation consists of:
// - meta data
// - translations for different target locales
// - parts of request information (NOTE: should not be used in the future)
type MultiTranslation struct {
	Source         Locale        `json:"from"`
	Query          string        `json:"phrase"`
	WiktionaryLink string        `json:"wiktionary-link"`
	Translations   []Translation `json:"meta-translations"`
}

// Translation contains information about translation to one locale.
type Translation struct {
	Target   Locale    `json:"dest"`
	WebURL   string    `json:"source-url"`
	Meanings []Meaning `json:"translations"`
}

// Meaning contains information about one meaning of translation.
type Meaning struct {
	Lexical        string   `json:"lexical"`
	TranslatedText string   `json:"phrase"`
	Sounds         []string `json:"sounds"`
	OriginName     string   `json:"source-name"`
	WebURL         string   `json:"source-url"`
}
