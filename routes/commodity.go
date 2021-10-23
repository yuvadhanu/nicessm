package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommodityRoutes : ""
func (route *Route) CommodityRoutes(r *mux.Router) {
	r.Handle("/commodity", Adapt(http.HandlerFunc(route.Handler.SaveCommodity))).Methods("POST")
	r.Handle("/commodity", Adapt(http.HandlerFunc(route.Handler.GetSingleCommodity))).Methods("GET")
	r.Handle("/commodity", Adapt(http.HandlerFunc(route.Handler.UpdateCommodity))).Methods("PUT")
	r.Handle("/commodity/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommodity))).Methods("PUT")
	r.Handle("/commodity/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommodity))).Methods("PUT")
	r.Handle("/commodity/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommodity))).Methods("DELETE")
	r.Handle("/commodity/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommodity))).Methods("POST")
	r.Handle("/commodity/insect/add", Adapt(http.HandlerFunc(route.Handler.AddInsectsCommodity))).Methods("POST")
	r.Handle("/commodity/disease/add", Adapt(http.HandlerFunc(route.Handler.AddDieseasesCommodity))).Methods("POST")
	r.Handle("/commodity/insect/delete", Adapt(http.HandlerFunc(route.Handler.DeleteInsectsCommodity))).Methods("POST")
	r.Handle("/commodity/disease/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDiseasesCommodity))).Methods("POST")
}
