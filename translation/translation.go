package translation

import (
	"github.com/eicca/translate-server/common"
)

// MultiReq contains translation request for set of locales.
type MultiReq struct {
	Source  common.Locale
	Targets []common.Locale
	Query   string
}

// Req contains translation request for one locale.
type Req struct {
	Source common.Locale
	Target common.Locale
	Query  string
}

// MultiTranslation consists of:
// - meta data
// - translations for different target locales
// - parts of request information (NOTE: should not be used in the future)
type MultiTranslation struct {
	Source         common.Locale `json:"from"`
	Query          string        `json:"phrase"`
	WiktionaryLink string        `json:"wiktionary-link"`
	Translations   []Translation `json:"meta-translations"`
}

// Translation contains information about translation to one locale.
type Translation struct {
	Target   common.Locale `json:"dest"`
	WebURL   string        `json:"source-url"`
	Meanings []Meaning     `json:"translations"`
}

// Meaning contains information about one meaning of translation.
type Meaning struct {
	Lexical        string   `json:"lexical"`
	TranslatedText string   `json:"phrase"`
	Sounds         []string `json:"sounds"`
	OriginName     string   `json:"source-name"`
	WebURL         string   `json:"source-url"`
}

func Translate(req MultiReq) (MultiTranslation, error) {

	return MultiTranslation{}, nil
}
