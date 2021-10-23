package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommodityCategoryRoutes : ""
func (route *Route) CommodityCategoryRoutes(r *mux.Router) {
	r.Handle("/commoditycategory", Adapt(http.HandlerFunc(route.Handler.SaveCommodityCategory))).Methods("POST")
	r.Handle("/commoditycategory", Adapt(http.HandlerFunc(route.Handler.GetSingleCommodityCategory))).Methods("GET")
	r.Handle("/commoditycategory", Adapt(http.HandlerFunc(route.Handler.UpdateCommodityCategory))).Methods("PUT")
	r.Handle("/commoditycategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommodityCategory))).Methods("PUT")
	r.Handle("/commoditycategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommodityCategory))).Methods("PUT")
	r.Handle("/commoditycategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommodityCategory))).Methods("DELETE")
	r.Handle("/commoditycategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommodityCategory))).Methods("POST")
}
