package gtranslate

import (
	"testing"

	"github.com/eicca/translate-server/common"
)

var detectTests = []struct {
	in  DetectReq
	out common.Locale
}{
	{
		DetectReq{
			Locales:        []common.Locale{"en", "de", "ru"},
			FallbackLocale: "de",
			Query:          "hallo",
		}, common.Locale("de"),
	},
	{
		DetectReq{
			Locales:        []common.Locale{"en", "ru"},
			FallbackLocale: "ru",
			Query:          "etwas",
		}, common.Locale("ru"),
	},
	{
		DetectReq{
			Locales:        []common.Locale{"en", "de", "ru"},
			FallbackLocale: "de",
		}, common.Locale("en"),
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
