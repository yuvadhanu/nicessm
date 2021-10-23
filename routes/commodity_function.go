package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommodityFunctionRoutes : ""
func (route *Route) CommodityFunctionRoutes(r *mux.Router) {
	r.Handle("/commodityfunction", Adapt(http.HandlerFunc(route.Handler.SaveCommodityFunction))).Methods("POST")
	r.Handle("/commodityfunction", Adapt(http.HandlerFunc(route.Handler.GetSingleCommodityFunction))).Methods("GET")
	r.Handle("/commodityfunction", Adapt(http.HandlerFunc(route.Handler.UpdateCommodityFunction))).Methods("PUT")
	r.Handle("/commodityfunction/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommodityFunction))).Methods("PUT")
	r.Handle("/commodityfunction/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommodityFunction))).Methods("PUT")
	r.Handle("/commodityfunction/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommodityFunction))).Methods("DELETE")
	r.Handle("/commodityfunction/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommodityFunction))).Methods("POST")
}
