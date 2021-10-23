package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//LiveStockVaccinationRoutes : ""
func (route *Route) LiveStockVaccinationRoutes(r *mux.Router) {
	r.Handle("/livestockvaccination", Adapt(http.HandlerFunc(route.Handler.SaveLiveStockVaccination))).Methods("POST")
	r.Handle("/livestockvaccination", Adapt(http.HandlerFunc(route.Handler.GetSingleLiveStockVaccination))).Methods("GET")
	r.Handle("/livestockvaccination", Adapt(http.HandlerFunc(route.Handler.UpdateLiveStockVaccination))).Methods("PUT")
	r.Handle("/livestockvaccination/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLiveStockVaccination))).Methods("PUT")
	r.Handle("/livestockvaccination/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLiveStockVaccination))).Methods("PUT")
	r.Handle("/livestockvaccination/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLiveStockVaccination))).Methods("DELETE")
	r.Handle("/livestockvaccination/filter", Adapt(http.HandlerFunc(route.Handler.FilterLiveStockVaccination))).Methods("POST")
}
