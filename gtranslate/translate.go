package gtranslate

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

// TranslateReq contains translation request information.
type TranslateReq struct {
	Source string
	Target string
	Query  string
}

// Locale is a two letters ISO locale.
type Locale string

type apiTranslateResp struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

// Translate translates a text from source locale to target locale.
func Translate(req TranslateReq) (string, error) {
	apiQuery, err := makeTranslateQuery(req)
	if err != nil {
		return "", err
	}

	data, err := get(apiQuery)
	if err != nil {
		return "", err
	}

	return parseTranslateResp(data)
}

func makeTranslateQuery(req TranslateReq) (*url.URL, error) {
	u, err := url.Parse(translateURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("source", req.Source)
	q.Set("target", req.Target)
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
