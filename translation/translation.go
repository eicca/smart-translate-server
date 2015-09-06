package translation

import (
	"fmt"

	"github.com/eicca/translate-server/data"
	"github.com/eicca/translate-server/gtranslate"
)

// Translate returns MultiTranslation filled from different sources.
func Translate(multiReq data.MultiTranslationReq) (data.MultiTranslation, error) {
	multiT := data.MultiTranslation{
		Source:         multiReq.Source,
		Query:          multiReq.Query,
		WiktionaryLink: wiktionaryLink(multiReq),
	}

	for _, target := range multiReq.Targets {
		req := data.TranslationReq{Source: multiReq.Source, Target: target, Query: multiReq.Query}
		translation, err := gtranslate.Translate(req)
		if err != nil {
			return data.MultiTranslation{}, err
		}

		multiT.Translations = append(multiT.Translations, translation)
	}

	return multiT, nil
}

func wiktionaryLink(multiReq data.MultiTranslationReq) string {
	return fmt.Sprintf("http://%s.wiktionary.org/wiki/%s", multiReq.Source, multiReq.Query)
}
