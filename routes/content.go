package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentRoutes(r *mux.Router) {
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.SaveContent))).Methods("POST")
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.GetSingleContent))).Methods("GET")
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.UpdateContent))).Methods("PUT")
	r.Handle("/content/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContent))).Methods("PUT")
	r.Handle("/content/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContent))).Methods("PUT")
	r.Handle("/content/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContent))).Methods("DELETE")
	r.Handle("/content/filter", Adapt(http.HandlerFunc(route.Handler.FilterContent))).Methods("POST")
}
