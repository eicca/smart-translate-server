package data

type SuggestionReq struct {
	Source Locale
	Target Locale
	Query  string
}

type Suggestion struct {
	Text   string `json:"phrase"`
	Locale Locale `json:"locale"`
}
