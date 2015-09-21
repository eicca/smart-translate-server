package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/eicca/translate-server/data"
	"github.com/eicca/translate-server/translation"
)

// ListenAndServeRest runs http server for REST API.
func ListenAndServeRest() {
	api := NewRest()
	log.Fatal(http.ListenAndServe(":3456", api.MakeHandler()))
}

// NewRest configures REST api handlers and wrappers.
func NewRest() *rest.Api {
	api := rest.NewApi()
	api.MakeHandler()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/translations", getTranslations),
		rest.Get("/suggestions", getSuggestions),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	return api
}

func getTranslations(w rest.ResponseWriter, r *rest.Request) {
	required := []string{"from", "dest-locales", "phrase"}
	params, err := getParams(r, required)
	if err != nil {
		rest.Error((w), err.Error(), 400)
		return
	}

	req := data.MultiTranslationReq{
		Source:  data.Locale(params["from"][0]),
		Targets: data.StringsAsLocales(params["dest-locales"]),
		Query:   params["phrase"][0],
	}
	multiT, err := translation.Translate(req)
	if err != nil {
		rest.Error((w), err.Error(), 500)
		return
	}
	w.WriteJson(&multiT)
}

func getSuggestions(w rest.ResponseWriter, r *rest.Request) {
	required := []string{"phrase", "locales", "fallback-locale"}
	params, err := getParams(r, required)
	if err != nil {
		rest.Error((w), err.Error(), 400)
		return
	}

	fmt.Println(params)
}

func getParams(r *rest.Request, params []string) (url.Values, error) {
	query := r.URL.Query()

	for _, param := range params {
		if len(query[param]) < 1 {
			return nil, fmt.Errorf("%s parameter is required", param)
		}
	}

	return query, nil
}