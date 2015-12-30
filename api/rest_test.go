package api

import (
	"fmt"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/eicca/translate-server/data"
)

func TestRequiredParamsValidation(t *testing.T) {
	payload := data.MultiTranslationReq{
		Query:   "hello",
		Targets: []data.Locale{"ru", "de"},
	}
	rec := makeRequest(t, "/translations", payload)
	rec.CodeIs(400)
}

func TestTranslations(t *testing.T) {
	payload := data.MultiTranslationReq{
		Query:   "hello",
		Targets: []data.Locale{"ru", "de"},
		Source:  data.Locale("en"),
	}
	rec := makeRequest(t, "/translations", payload)
	rec.CodeIs(200)

	var mt data.MultiTranslation
	unmarshal(t, rec, &mt)

	if mt.Translations == nil {
		t.Error("Expected to return translations")
	}
}

func TestSuggestions(t *testing.T) {
	payload := data.SuggestionReq{
		Query:          "irgend",
		Locales:        []data.Locale{"en", "de"},
		FallbackLocale: data.Locale("en"),
	}
	rec := makeRequest(t, "/suggestions", payload)
	rec.CodeIs(200)

	var ss []data.Suggestion
	unmarshal(t, rec, &ss)

	if ss == nil {
		t.Error("Expected to return suggestions")
	}
}

func TestHealth(t *testing.T) {
	rec := makeRequest(t, "/health", nil)
	rec.CodeIs(200)
}

func makeRequest(t *testing.T, path string, payload interface{}) *test.Recorded {
	api := NewRest()
	req := test.MakeSimpleRequest("POST", fmt.Sprintf("http://1.2.3.4%s", path), payload)
	return test.RunRequest(t, api.MakeHandler(), req)
}

func unmarshal(t *testing.T, rec *test.Recorded, in interface{}) {
	if err := rec.DecodeJsonPayload(in); err != nil {
		t.Fatal(err)
	}
}
