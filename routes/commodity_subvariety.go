package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommoditySubVarietyRoutes : ""
func (route *Route) CommoditySubVarietyRoutes(r *mux.Router) {
	r.Handle("/commoditysubvariety", Adapt(http.HandlerFunc(route.Handler.SaveCommoditySubVariety))).Methods("POST")
	r.Handle("/commoditysubvariety", Adapt(http.HandlerFunc(route.Handler.GetSingleCommoditySubVariety))).Methods("GET")
	r.Handle("/commoditysubvariety", Adapt(http.HandlerFunc(route.Handler.UpdateCommoditySubVariety))).Methods("PUT")
	r.Handle("/commoditysubvariety/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommoditySubVariety))).Methods("PUT")
	r.Handle("/commoditysubvariety/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommoditySubVariety))).Methods("PUT")
	r.Handle("/commoditysubvariety/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommoditySubVariety))).Methods("DELETE")
	r.Handle("/commoditysubvariety/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommoditySubVariety))).Methods("POST")
}
