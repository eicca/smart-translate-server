package glosbe

import (
	"testing"

	"github.com/eicca/translate-server/data"
)

var suggestTests = []struct {
	in  data.SuggestionReq
	out []data.Suggestion
}{
	{
		data.SuggestionReq{
			Locales:        []data.Locale{"en", "de", "ru"},
			FallbackLocale: data.Locale("ru"),
			Query:          "irgend",
		},
		[]data.Suggestion{
			data.Suggestion{Text: "irgendein", Locale: "de"},
			data.Suggestion{Text: "irgend etwas", Locale: "de"},
		},
	},
	{
		data.SuggestionReq{
			Locales:        []data.Locale{"en", "ru"},
			FallbackLocale: data.Locale("ru"),
			Query:          "irgend",
		},
		[]data.Suggestion{
			data.Suggestion{Text: "Iris", Locale: "ru"},
		},
	},
}

func TestSuggest(t *testing.T) {
	for _, tt := range suggestTests {
		resMap := make(map[string]data.Locale)

		res, err := Suggest(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		for _, s := range res {
			resMap[s.Text] = s.Locale
		}

		for _, s := range tt.out {
			locale, ok := resMap[s.Text]

			if !ok || s.Locale != locale {
				t.Errorf("Expected TestSuggest to return a slice containing %+v, got slice of %+v", s, res)
			}
		}
	}
}
