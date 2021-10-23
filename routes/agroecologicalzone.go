package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) AgroEcologicalZoneRoutes(r *mux.Router) {
	r.Handle("/agroEcologicalZone", Adapt(http.HandlerFunc(route.Handler.SaveAgroEcologicalZone))).Methods("POST")
	r.Handle("/agroEcologicalZone", Adapt(http.HandlerFunc(route.Handler.GetSingleAgroEcologicalZone))).Methods("GET")
	r.Handle("/agroEcologicalZone", Adapt(http.HandlerFunc(route.Handler.UpdateAgroEcologicalZone))).Methods("PUT")
	r.Handle("/agroEcologicalZone/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAgroEcologicalZone))).Methods("PUT")
	r.Handle("/agroEcologicalZone/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAgroEcologicalZone))).Methods("PUT")
	r.Handle("/agroEcologicalZone/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAgroEcologicalZone))).Methods("DELETE")
	r.Handle("/agroEcologicalZone/filter", Adapt(http.HandlerFunc(route.Handler.FilterAgroEcologicalZone))).Methods("POST")
}
