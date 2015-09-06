package gtranslate

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/eicca/translate-server/data"
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
func Translate(req data.TranslationReq) (data.Translation, error) {
	apiQuery, err := makeTranslateQuery(req)
	if err != nil {
		return data.Translation{}, err
	}

	rawData, err := get(apiQuery)
	if err != nil {
		return data.Translation{}, err
	}

	translatedText, err := parseTranslateResp(rawData)
	if err != nil {
		return data.Translation{}, err
	}

	return translation(req, translatedText), nil
}

func translation(req data.TranslationReq, translatedText string) data.Translation {
	translationWebURL := fmt.Sprintf(
		"%s/#%s/%s/%s",
		gtranslateWebURL, req.Source, req.Target, req.Query,
	)

	// gtranslate returns always only one meaning.
	meaning := meaning(req, translatedText)

	return data.Translation{
		Target:   req.Target,
		WebURL:   translationWebURL,
		Meanings: []data.Meaning{meaning},
	}
}

func meaning(req data.TranslationReq, translatedText string) data.Meaning {
	// meaningWebURL is a reverse of translation.
	meaningWebURL := fmt.Sprintf(
		"%s/#%s/%s/%s",
		gtranslateWebURL, req.Target, req.Source, translatedText,
	)

	return data.Meaning{
		TranslatedText: translatedText,
		OriginName:     originName,
		WebURL:         meaningWebURL,
	}
}

func makeTranslateQuery(req data.TranslationReq) (*url.URL, error) {
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

func parseTranslateResp(rawData []byte) (string, error) {
	apiResp := apiTranslateResp{}
	if err := json.Unmarshal(rawData, &apiResp); err != nil {
		return "", fmt.Errorf("Error during google Translate API unmarshaling: %s", err)
	}

	translations := apiResp.Data.Translations
	if len(translations) < 1 {
		return "", fmt.Errorf("No data was returned from google translate API. Response: %s \n Unmarshaled response: %v", rawData, apiResp)
	}

	text := translations[0].TranslatedText
	return text, nil
}
