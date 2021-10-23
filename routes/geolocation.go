package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StateRoutes : ""
func (route *Route) StateRoutes(r *mux.Router) {
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.SaveState))).Methods("POST")
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.GetSingleState))).Methods("GET")
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.UpdateState))).Methods("PUT")
	r.Handle("/state/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableState))).Methods("PUT")
	r.Handle("/state/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableState))).Methods("PUT")
	r.Handle("/state/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteState))).Methods("DELETE")
	r.Handle("/state/filter", Adapt(http.HandlerFunc(route.Handler.FilterState))).Methods("POST")
}

//DistrictRoutes : ""
func (route *Route) DistrictRoutes(r *mux.Router) {
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.SaveDistrict))).Methods("POST")
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrict))).Methods("GET")
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.UpdateDistrict))).Methods("PUT")
	r.Handle("/district/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrict))).Methods("PUT")
	r.Handle("/district/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrict))).Methods("PUT")
	r.Handle("/district/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrict))).Methods("DELETE")
	r.Handle("/district/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrict))).Methods("POST")
}

//BlockRoutes : ""
func (route *Route) BlockRoutes(r *mux.Router) {
	r.Handle("/block", Adapt(http.HandlerFunc(route.Handler.SaveBlock))).Methods("POST")
	r.Handle("/block", Adapt(http.HandlerFunc(route.Handler.GetSingleBlock))).Methods("GET")
	r.Handle("/block", Adapt(http.HandlerFunc(route.Handler.UpdateBlock))).Methods("PUT")
	r.Handle("/block/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBlock))).Methods("PUT")
	r.Handle("/block/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBlock))).Methods("PUT")
	r.Handle("/block/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBlock))).Methods("DELETE")
	r.Handle("/block/filter", Adapt(http.HandlerFunc(route.Handler.FilterBlock))).Methods("POST")
}

//GramPanchayatRoutes : ""
func (route *Route) GramPanchayatRoutes(r *mux.Router) {
	r.Handle("/grampanchayat", Adapt(http.HandlerFunc(route.Handler.SaveGramPanchayat))).Methods("POST")
	r.Handle("/grampanchayat", Adapt(http.HandlerFunc(route.Handler.GetSingleGramPanchayat))).Methods("GET")
	r.Handle("/grampanchayat", Adapt(http.HandlerFunc(route.Handler.UpdateGramPanchayat))).Methods("PUT")
	r.Handle("/grampanchayat/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableGramPanchayat))).Methods("PUT")
	r.Handle("/grampanchayat/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableGramPanchayat))).Methods("PUT")
	r.Handle("/grampanchayat/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteGramPanchayat))).Methods("DELETE")
	r.Handle("/grampanchayat/filter", Adapt(http.HandlerFunc(route.Handler.FilterGramPanchayat))).Methods("POST")
}

//VillageRoutes : ""
func (route *Route) VillageRoutes(r *mux.Router) {
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.SaveVillage))).Methods("POST")
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.GetSingleVillage))).Methods("GET")
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.UpdateVillage))).Methods("PUT")
	r.Handle("/village/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVillage))).Methods("PUT")
	r.Handle("/village/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVillage))).Methods("PUT")
	r.Handle("/village/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVillage))).Methods("DELETE")
	r.Handle("/village/filter", Adapt(http.HandlerFunc(route.Handler.FilterVillage))).Methods("POST")
}
