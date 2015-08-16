package gtranslate

import (
	"testing"
)

var detectTests = []struct {
	in  DetectReq
	out Locale
}{
	{
		DetectReq{
			Locales:        []Locale{"en", "de", "ru"},
			FallbackLocale: "de",
			Query:          "hallo",
		}, Locale("de"),
	},
	{
		DetectReq{
			Locales:        []Locale{"en", "ru"},
			FallbackLocale: "ru",
			Query:          "etwas",
		}, Locale("ru"),
	},
	{
		DetectReq{
			Locales:        []Locale{"en", "de", "ru"},
			FallbackLocale: "de",
		}, Locale("en"),
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
