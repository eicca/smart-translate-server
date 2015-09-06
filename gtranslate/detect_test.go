package gtranslate

import (
	"testing"

	"github.com/eicca/translate-server/data"
)

var detectTests = []struct {
	in  DetectReq
	out data.Locale
}{
	{
		DetectReq{
			Locales:        []data.Locale{"en", "de", "ru"},
			FallbackLocale: "de",
			Query:          "hallo",
		}, data.Locale("de"),
	},
	{
		DetectReq{
			Locales:        []data.Locale{"en", "ru"},
			FallbackLocale: "ru",
			Query:          "etwas",
		}, data.Locale("ru"),
	},
	{
		DetectReq{
			Locales:        []data.Locale{"en", "de", "ru"},
			FallbackLocale: "de",
		}, data.Locale("en"),
	},
}

func TestDetect(t *testing.T) {
	for _, tt := range detectTests {
		res, err := Detect(tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if res != tt.out {
			t.Errorf("Wrong detection. Expected to detect %s for %+v got %s ",
				tt.out, tt.in, res)
		}
	}
}
