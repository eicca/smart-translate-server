package gtranslate

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	t "github.com/eicca/translate-server/translation"
)

const (
	gtranslateWebURL = "https://translate.google.com"
	originName       = "google"
)

type apiTranslateResp struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

// Translate translates a text from source locale to target locale.
func Translate(req t.Req) (t.Translation, error) {
	apiQuery, err := makeTranslateQuery(req)
	if err != nil {
		return t.Translation{}, err
	}

	data, err := get(apiQuery)
	if err != nil {
		return t.Translation{}, err
	}

	translatedText, err := parseTranslateResp(data)
	if err != nil {
		return t.Translation{}, err
	}

	return translation(req, translatedText), nil
}

func translation(req t.Req, translatedText string) t.Translation {
	translationWebURL := fmt.Sprintf(
		"%s/#%s/%s/%s",
		gtranslateWebURL, req.Source, req.Target, req.Query,
	)

	// gtranslate returns always only one meaning.
	meaning := meaning(req, translatedText)

	return t.Translation{
		Target:   req.Target,
		WebURL:   translationWebURL,
		Meanings: []t.Meaning{meaning},
	}
}

func meaning(req t.Req, translatedText string) t.Meaning {
	// meaningWebURL is a reverse of translation.
	meaningWebURL := fmt.Sprintf(
		"%s/#%s/%s/%s",
		gtranslateWebURL, req.Target, req.Source, translatedText,
	)

	return t.Meaning{
		TranslatedText: translatedText,
		OriginName:     originName,
		WebURL:         meaningWebURL,
	}
}

func makeTranslateQuery(req t.Req) (*url.URL, error) {
	u, err := url.Parse(translateURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("source", string(req.Source))
	q.Set("target", string(req.Target))
	q.Set("q", req.Query)
	q.Set("key", os.Getenv("GOOGLE_API_KEY"))
	u.RawQuery = q.Encode()
	return u, nil
}

func parseTranslateResp(data []byte) (string, error) {
	apiResp := apiTranslateResp{}
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return "", fmt.Errorf("Error during google Translate API unmarshaling: %s", err)
	}

	translations := apiResp.Data.Translations
	if len(translations) < 1 {
		return "", fmt.Errorf("No data was returned from google translate API. Response: %s \n Unmarshaled response: %v", data, apiResp)
	}

	text := translations[0].TranslatedText
	return text, nil
}
