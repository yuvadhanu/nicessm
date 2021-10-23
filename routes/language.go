package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) LanguageRoutes(r *mux.Router) {
	r.Handle("/langauage", Adapt(http.HandlerFunc(route.Handler.SaveLanguage))).Methods("POST")
	r.Handle("/langauage", Adapt(http.HandlerFunc(route.Handler.GetSingleLanguage))).Methods("GET")
	r.Handle("/langauage", Adapt(http.HandlerFunc(route.Handler.UpdateLanguage))).Methods("PUT")
	r.Handle("/langauage/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLanguage))).Methods("PUT")
	r.Handle("/langauage/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLanguage))).Methods("PUT")
	r.Handle("/langauage/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLanguage))).Methods("DELETE")
	r.Handle("/langauage/filter", Adapt(http.HandlerFunc(route.Handler.FilterLanguage))).Methods("POST")
}
