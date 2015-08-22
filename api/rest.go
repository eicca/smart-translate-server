package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
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
	query := r.URL.Query()

	for _, param := range required {
		if len(query[param]) < 1 {
			rest.Error((w), fmt.Sprintf("%s parameter is required", param), 400)
			return
		}
	}
}

func getSuggestions(w rest.ResponseWriter, r *rest.Request) {
	log.Println(r)
}
