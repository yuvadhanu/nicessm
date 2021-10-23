package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SoilTypeRoutes : ""
func (route *Route) SoilTypeRoutes(r *mux.Router) {
	r.Handle("/soiltype", Adapt(http.HandlerFunc(route.Handler.SaveSoilType))).Methods("POST")
	r.Handle("/soiltype", Adapt(http.HandlerFunc(route.Handler.GetSingleSoilType))).Methods("GET")
	r.Handle("/soiltype", Adapt(http.HandlerFunc(route.Handler.UpdateSoilType))).Methods("PUT")
	r.Handle("/soiltype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSoilType))).Methods("PUT")
	r.Handle("/soiltype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSoilType))).Methods("PUT")
	r.Handle("/soiltype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSoilType))).Methods("DELETE")
	r.Handle("/soiltype/filter", Adapt(http.HandlerFunc(route.Handler.FilterSoilType))).Methods("POST")
}
