package glosbe

import (
	"github.com/eicca/translate-server/data"
)

// Translate translates a text from source locale to target locale.
func Translate(req data.TranslationReq) (data.Translation, error) {
	t := data.Translation{}
	m := data.Meaning{}
	m.TranslatedText = "translated text"
	t.Meanings = []data.Meaning{m}
	t.Target = req.Target

	return t, nil
}
