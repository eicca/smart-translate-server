package gtranslate

import (
	"reflect"
	"testing"

	"github.com/eicca/translate-server/common"
	tr "github.com/eicca/translate-server/translation"
)

func TestTranslateText(t *testing.T) {
	req := tr.Req{Source: "en", Target: "de", Query: "hello"}
	expectedRes := tr.Translation{
		Target: common.Locale("de"),
		WebURL: "https://translate.google.com/#en/de/hello",
		Meanings: []tr.Meaning{tr.Meaning{
			TranslatedText: "Hallo",
			OriginName:     "google",
			WebURL:         "https://translate.google.com/#de/en/Hallo",
		}},
	}
	res, err := Translate(req)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res, expectedRes) {
		t.Errorf("Wrong translation for '%s'.\nExpected: %+v\nGot: %+v",
			req.Query, expectedRes, res)
	}
}

func TestTranslateApiError(t *testing.T) {
	req := tr.Req{Source: "en", Query: "Hello"}
	_, err := Translate(req)
	if err == nil {
		t.Fatal("Translate haven't returned error on invalid request.")
	}
}
