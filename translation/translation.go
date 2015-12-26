package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

		// Try to translate with glosbe first.
		translation, err := glosbeTranslate(req)
		if err != nil {
			return data.MultiTranslation{}, fmt.Errorf("glosbe translation failed: %s", err)
		}

		// If there's no glosbe translation - fallback to google translation
		if len(translation.Meanings) < 1 {
			translation, err = gtranslate.Translate(req)
			if err != nil {
				return data.MultiTranslation{}, fmt.Errorf("google translation failed: %s", err)
			}
		}

		multiT.Translations = append(multiT.Translations, translation)
	}

	return multiT, nil
}

func wiktionaryLink(multiReq data.MultiTranslationReq) string {
	return fmt.Sprintf("http://%s.wiktionary.org/wiki/%s", multiReq.Source, multiReq.Query)
}

func glosbeTranslate(req data.TranslationReq) (data.Translation, error) {
	host := os.Getenv("GLOSBE_TRANSLATE_HOST")
	url := "http://" + host + "/translations"
	body, err := json.Marshal(req)
	if err != nil {
		return data.Translation{}, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return data.Translation{}, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return data.Translation{}, err
	}

	var t data.Translation
	err = json.Unmarshal(body, &t)
	return t, err
}
