package gtranslate

import (
	"testing"
)

func TestTranslateText(t *testing.T) {
	req := TranslateReq{Source: "en", Target: "de", Query: "hello"}
	expectedRes := "Hallo"
	res, err := Translate(req)
	if err != nil {
		t.Fatal(err)
	}

	if res != expectedRes {
		t.Errorf("Wrong translation. Expected to translate %s to %s got %s ",
			req.Query, expectedRes, res)
	}
}

func TestTranslateApiError(t *testing.T) {
	req := TranslateReq{Source: "en", Query: "Hello"}
	_, err := Translate(req)
	if err == nil {
		t.Fatal("Translate haven't returned error on invalid request.")
	}
}
