package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommodityStageRoutes : ""
func (route *Route) CommodityStageRoutes(r *mux.Router) {
	r.Handle("/commoditystage", Adapt(http.HandlerFunc(route.Handler.SaveCommodityStage))).Methods("POST")
	r.Handle("/commoditystage", Adapt(http.HandlerFunc(route.Handler.GetSingleCommodityStage))).Methods("GET")
	r.Handle("/commoditystage", Adapt(http.HandlerFunc(route.Handler.UpdateCommodityStage))).Methods("PUT")
	r.Handle("/commoditystage/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommodityStage))).Methods("PUT")
	r.Handle("/commoditystage/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommodityStage))).Methods("PUT")
	r.Handle("/commoditystage/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommodityStage))).Methods("DELETE")
	r.Handle("/commoditystage/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommodityStage))).Methods("POST")
}
