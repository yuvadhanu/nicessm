package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommodityVarietyRoutes : ""
func (route *Route) CommodityVarietyRoutes(r *mux.Router) {
	r.Handle("/commodityvariety", Adapt(http.HandlerFunc(route.Handler.SaveCommodityVariety))).Methods("POST")
	r.Handle("/commodityvariety", Adapt(http.HandlerFunc(route.Handler.GetSingleCommodityVariety))).Methods("GET")
	r.Handle("/commodityvariety", Adapt(http.HandlerFunc(route.Handler.UpdateCommodityVariety))).Methods("PUT")
	r.Handle("/commodityvariety/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommodityVariety))).Methods("PUT")
	r.Handle("/commodityvariety/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommodityVariety))).Methods("PUT")
	r.Handle("/commodityvariety/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommodityVariety))).Methods("DELETE")
	r.Handle("/commodityvariety/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommodityVariety))).Methods("POST")
}
