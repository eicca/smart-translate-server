package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/eicca/translate-server/data"
	"github.com/eicca/translate-server/glosbe"
	"github.com/eicca/translate-server/translation"
)

// ListenAndServeRest runs http server for REST API.
func ListenAndServeRest(port string) {
	api := NewRest()
	log.Fatal(http.ListenAndServe(port, api.MakeHandler()))
}

// NewRest configures REST api handlers and wrappers.
func NewRest() *rest.Api {
	api := rest.NewApi()
	api.MakeHandler()
	api.Use(rest.DefaultDevStack...)

	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

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

	req := data.SuggestionReq{
		Locales:        data.StringsAsLocales(params["locales"]),
		FallbackLocale: data.Locale(params["fallback-locale"][0]),
		Query:          params["phrase"][0],
	}

	suggestions, err := glosbe.Suggest(req)
	if err != nil {
		rest.Error((w), err.Error(), 500)
		return
	}
	w.WriteJson(&suggestions)
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
