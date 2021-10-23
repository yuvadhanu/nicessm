package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectFarmerRoutes : ""
func (route *Route) ProjectFarmerRoutes(r *mux.Router) {
	r.Handle("/projectfarmer", Adapt(http.HandlerFunc(route.Handler.SaveProjectFarmer))).Methods("POST")
	r.Handle("/projectfarmer", Adapt(http.HandlerFunc(route.Handler.GetSingleProjectFarmer))).Methods("GET")
	r.Handle("/projectfarmer", Adapt(http.HandlerFunc(route.Handler.UpdateProjectFarmer))).Methods("PUT")
	r.Handle("/projectfarmer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProjectFarmer))).Methods("PUT")
	r.Handle("/projectfarmer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProjectFarmer))).Methods("PUT")
	r.Handle("/projectfarmer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProjectFarmer))).Methods("DELETE")
	r.Handle("/projectfarmer/filter", Adapt(http.HandlerFunc(route.Handler.FilterProjectFarmer))).Methods("POST")
}
