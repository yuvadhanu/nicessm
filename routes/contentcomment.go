package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentCommentRoutes(r *mux.Router) {
	r.Handle("/contentcomment", Adapt(http.HandlerFunc(route.Handler.SaveContentComment))).Methods("POST")
	r.Handle("/contentcomment", Adapt(http.HandlerFunc(route.Handler.GetSingleContentComment))).Methods("GET")
	r.Handle("/contentcomment", Adapt(http.HandlerFunc(route.Handler.UpdateContentComment))).Methods("PUT")
	r.Handle("/contentcomment/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContentComment))).Methods("PUT")
	r.Handle("/contentcomment/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContentComment))).Methods("PUT")
	r.Handle("/contentcomment/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContentComment))).Methods("DELETE")
	r.Handle("/contentcomment/filter", Adapt(http.HandlerFunc(route.Handler.FilterContentComment))).Methods("POST")
}
