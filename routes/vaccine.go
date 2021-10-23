package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) VaccineRoutes(r *mux.Router) {
	r.Handle("/vaccine", Adapt(http.HandlerFunc(route.Handler.SaveVaccine))).Methods("POST")
	r.Handle("/vaccine", Adapt(http.HandlerFunc(route.Handler.GetSingleVaccine))).Methods("GET")
	r.Handle("/vaccine", Adapt(http.HandlerFunc(route.Handler.UpdateVaccine))).Methods("PUT")
	r.Handle("/vaccine/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVaccine))).Methods("PUT")
	r.Handle("/vaccine/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVaccine))).Methods("PUT")
	r.Handle("/vaccine/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVaccine))).Methods("DELETE")
	r.Handle("/vaccine/filter", Adapt(http.HandlerFunc(route.Handler.FilterVaccine))).Methods("POST")
}
