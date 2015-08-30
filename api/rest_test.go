package api

import (
	"fmt"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
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
}

func TestSuggestions(t *testing.T) {
	path := "/suggestions?phrase=hello&locales=en&locales=de&fallback-locale=de"
	rec := makeRequest(t, path)
	rec.CodeIs(200)
}

func makeRequest(t *testing.T, path string) *test.Recorded {
	api := NewRest()
	req := test.MakeSimpleRequest("GET", fmt.Sprintf("http://1.2.3.4%s", path), nil)
	return test.RunRequest(t, api.MakeHandler(), req)
}
