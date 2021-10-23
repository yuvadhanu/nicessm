package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectStateRoutes : ""
func (route *Route) ProjectStateRoutes(r *mux.Router) {
	r.Handle("/projectstate", Adapt(http.HandlerFunc(route.Handler.SaveProjectState))).Methods("POST")
	r.Handle("/projectstate", Adapt(http.HandlerFunc(route.Handler.GetSingleProjectState))).Methods("GET")
	r.Handle("/projectstate", Adapt(http.HandlerFunc(route.Handler.UpdateProjectState))).Methods("PUT")
	r.Handle("/projectstate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProjectState))).Methods("PUT")
	r.Handle("/projectstate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProjectState))).Methods("PUT")
	r.Handle("/projectstate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProjectState))).Methods("DELETE")
	r.Handle("/projectstate/filter", Adapt(http.HandlerFunc(route.Handler.FilterProjectState))).Methods("POST")
}
