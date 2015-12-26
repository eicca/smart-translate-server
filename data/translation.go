package data

// MultiTranslationReq contains translation request for set of locales.
type MultiTranslationReq struct {
	Source  Locale   `json:"source"`
	Targets []Locale `json:"targets"`
	Query   string   `json:"query"`
}

// TranslationReq contains translation request for one locale.
type TranslationReq struct {
	Source Locale `json:"source"`
	Target Locale `json:"target"`
	Query  string `json:"query"`
}

// MultiTranslation consists of:
// - meta data
// - translations for different target locales
// - parts of request information (NOTE: should not be used in the future)
type MultiTranslation struct {
	Source         Locale        `json:"source"`
	Query          string        `json:"query"`
	WiktionaryLink string        `json:"wiktionary-link"`
	Translations   []Translation `json:"translations"`
}

// Translation contains information about translation to one locale.
type Translation struct {
	Target   Locale    `json:"target"`
	WebURL   string    `json:"web-url"`
	Meanings []Meaning `json:"meanings"`
}

// Meaning contains information about one meaning of translation.
type Meaning struct {
	Lexical        []string `json:"lexical"`
	TranslatedText string   `json:"translated-text"`
	Sounds         []string `json:"sounds"`
	OriginName     string   `json:"origin-name"`
	WebURL         string   `json:"web-url"`
}
