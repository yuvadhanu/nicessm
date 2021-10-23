package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentTranslationRoutes(r *mux.Router) {
	r.Handle("/contenttranslation", Adapt(http.HandlerFunc(route.Handler.SaveContentTranslation))).Methods("POST")
	r.Handle("/contenttranslation", Adapt(http.HandlerFunc(route.Handler.GetSingleContentTranslation))).Methods("GET")
	r.Handle("/contenttranslation", Adapt(http.HandlerFunc(route.Handler.UpdateContentTranslation))).Methods("PUT")
	r.Handle("/contenttranslation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContentTranslation))).Methods("PUT")
	r.Handle("/contenttranslation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContentTranslation))).Methods("PUT")
	r.Handle("/contenttranslation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContentTranslation))).Methods("DELETE")
	r.Handle("/contenttranslation/filter", Adapt(http.HandlerFunc(route.Handler.FilterContentTranslation))).Methods("POST")
}
