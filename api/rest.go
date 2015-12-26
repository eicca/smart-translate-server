package api

import (
	"log"
	"net/http"

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
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	router, err := rest.MakeRouter(
		rest.Post("/translations", translations),
		rest.Post("/suggestions", suggestions),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	return api
}

func translations(w rest.ResponseWriter, r *rest.Request) {
	var req data.MultiTranslationReq
	err := r.DecodeJsonPayload(&req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Targets == nil {
		rest.Error(w, "target locales required", 400)
		return
	}
	if req.Source == "" {
		rest.Error(w, "source locale required", 400)
		return
	}

	multiT, err := translation.Translate(req)
	if err != nil {
		rest.Error((w), err.Error(), 500)
		return
	}
	w.WriteJson(&multiT)
}

func suggestions(w rest.ResponseWriter, r *rest.Request) {
	var req data.SuggestionReq
	err := r.DecodeJsonPayload(&req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.FallbackLocale == "" {
		req.FallbackLocale = data.Locale("en")
	}
	if req.Locales == nil {
		rest.Error(w, "locales required", 400)
		return
	}

	suggestions, err := glosbe.Suggest(req)
	if err != nil {
		rest.Error((w), err.Error(), 500)
		return
	}
	w.WriteJson(&suggestions)
}
