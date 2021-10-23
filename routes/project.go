package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectRoutes : ""
func (route *Route) ProjectRoutes(r *mux.Router) {
	r.Handle("/project", Adapt(http.HandlerFunc(route.Handler.SaveProject))).Methods("POST")
	r.Handle("/project", Adapt(http.HandlerFunc(route.Handler.GetSingleProject))).Methods("GET")
	r.Handle("/project", Adapt(http.HandlerFunc(route.Handler.UpdateProject))).Methods("PUT")
	r.Handle("/project/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProject))).Methods("PUT")
	r.Handle("/project/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProject))).Methods("PUT")
	r.Handle("/project/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProject))).Methods("DELETE")
	r.Handle("/project/filter", Adapt(http.HandlerFunc(route.Handler.FilterProject))).Methods("POST")
}
