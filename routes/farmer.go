package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerRoutes : ""
func (route *Route) FarmerRoutes(r *mux.Router) {
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.SaveFarmer))).Methods("POST")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmer))).Methods("GET")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.UpdateFarmer))).Methods("PUT")
	r.Handle("/farmer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmer))).Methods("DELETE")
	r.Handle("/farmer/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmer))).Methods("POST")
}
