package api

import (
	"fmt"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/eicca/translate-server/data"
)

func TestRequiredParamsValidation(t *testing.T) {
	path := "/translations?dest-locales=ru&dest-locales=de&phrase=hello"
	rec := makeRequest(t, path)
	rec.CodeIs(400)
}

func TestTranslations(t *testing.T) {
	path := "/translations?from=en&dest-locales=ru&dest-locales=de&phrase=hello"
	rec := makeRequest(t, path)
	rec.CodeIs(200)

	var mt data.MultiTranslation
	unmarshal(t, rec, &mt)

	if mt.Translations == nil {
		t.Error("Expected to return translations")
	}
}

func TestSuggestions(t *testing.T) {
	path := "/suggestions?phrase=irgend&locales=en&locales=de&fallback-locale=en"
	rec := makeRequest(t, path)
	rec.CodeIs(200)

	var ss []data.Suggestion
	unmarshal(t, rec, &ss)

	if ss == nil {
		t.Error("Expected to return suggestions")
	}
}

func makeRequest(t *testing.T, path string) *test.Recorded {
	api := NewRest()
	req := test.MakeSimpleRequest("GET", fmt.Sprintf("http://1.2.3.4%s", path), nil)
	return test.RunRequest(t, api.MakeHandler(), req)
}

func unmarshal(t *testing.T, rec *test.Recorded, in interface{}) {
	if err := rec.DecodeJsonPayload(in); err != nil {
		t.Fatal(err)
	}
}
