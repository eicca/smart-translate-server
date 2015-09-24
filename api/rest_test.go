package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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
	compareResponse(t, rec, "hello|en->ru,de")
}

func TestSuggestions(t *testing.T) {
	path := "/suggestions?phrase=irgend&locales=en&locales=de&fallback-locale=en"
	rec := makeRequest(t, path)
	rec.CodeIs(200)
	compareResponse(t, rec, "irgend|en,de")
}

func makeRequest(t *testing.T, path string) *test.Recorded {
	api := NewRest()
	req := test.MakeSimpleRequest("GET", fmt.Sprintf("http://1.2.3.4%s", path), nil)
	return test.RunRequest(t, api.MakeHandler(), req)
}

func compareResponse(t *testing.T, rec *test.Recorded, fileName string) {
	expected := exampleRequest(t, fileName)

	var actual interface{}
	fmt.Println(rec.Recorder.Body)
	if err := rec.DecodeJsonPayload(&actual); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong response for '%s' request.\nExpected: %+v\nGot: %+v",
			fileName, expected, actual)
	}
}

func exampleRequest(t *testing.T, fileName string) interface{} {
	dat, err := ioutil.ReadFile("example_requests/" + fileName + ".json")
	if err != nil {
		t.Fatal(err)
	}

	var expected interface{}
	if err = json.Unmarshal(dat, &expected); err != nil {
		t.Fatal(err)
	}

	return expected
}
