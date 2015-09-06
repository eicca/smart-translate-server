package gtranslate

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/eicca/translate-server/data"
)

var (
	detectURL = fmt.Sprintf("%s/detect", translateURL)
)

// DetectReq contains translation request information.
type DetectReq struct {
	Locales        []data.Locale
	FallbackLocale data.Locale
	Query          string
}

type apiDetectResp struct {
	Data struct {
		Detections [][]struct {
			Language data.Locale `json:"language"`
		} `json:"detections"`
	} `json:"data"`
}

// Detect detects locale of the text.
func Detect(req DetectReq) (data.Locale, error) {
	apiQuery, err := makeDetectQuery(req)
	if err != nil {
		return "", err
	}

	data, err := get(apiQuery)
	if err != nil {
		return "", err
	}

	locale, err := parseDetectResp(data)
	if err != nil {
		return "", err
	}

	return normalizeLocale(locale, req), nil
}

func makeDetectQuery(req DetectReq) (*url.URL, error) {
	u, err := url.Parse(detectURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("q", req.Query)
	q.Set("key", os.Getenv("GOOGLE_API_KEY"))
	u.RawQuery = q.Encode()
	return u, nil
}

func parseDetectResp(data []byte) (data.Locale, error) {
	apiResp := apiDetectResp{}
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return "", fmt.Errorf("Error during google Translate API unmarshaling: %s", err)
	}

	detections := apiResp.Data.Detections
	if len(detections) < 1 || len(detections[0]) < 1 {
		return "", fmt.Errorf("No data was returned from google translate API. Response: %s \n Unmarshaled response: %v", data, apiResp)
	}

	locale := detections[0][0].Language
	return locale, nil
}

func normalizeLocale(locale data.Locale, req DetectReq) data.Locale {
	for _, reqLocale := range req.Locales {
		if reqLocale == locale {
			return locale
		}
	}
	return req.FallbackLocale
}
