package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectPartnerRoutes : ""
func (route *Route) ProjectPartnerRoutes(r *mux.Router) {
	r.Handle("/projectpartner", Adapt(http.HandlerFunc(route.Handler.SaveProjectPartner))).Methods("POST")
	r.Handle("/projectpartner", Adapt(http.HandlerFunc(route.Handler.GetSingleProjectPartner))).Methods("GET")
	r.Handle("/projectpartner", Adapt(http.HandlerFunc(route.Handler.UpdateProjectPartner))).Methods("PUT")
	r.Handle("/projectpartner/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProjectPartner))).Methods("PUT")
	r.Handle("/projectpartner/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProjectPartner))).Methods("PUT")
	r.Handle("/projectpartner/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProjectPartner))).Methods("DELETE")
	r.Handle("/projectpartner/filter", Adapt(http.HandlerFunc(route.Handler.FilterProjectPartner))).Methods("POST")
}
