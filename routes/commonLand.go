package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommonLandRoutes : ""
func (route *Route) CommonLandRoutes(r *mux.Router) {
	r.Handle("/commonland", Adapt(http.HandlerFunc(route.Handler.SaveCommonLand))).Methods("POST")
	r.Handle("/commonland", Adapt(http.HandlerFunc(route.Handler.GetSingleCommonLand))).Methods("GET")
	r.Handle("/commonland", Adapt(http.HandlerFunc(route.Handler.UpdateCommonLand))).Methods("PUT")
	r.Handle("/commonland/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommonLand))).Methods("PUT")
	r.Handle("/commonland/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommonLand))).Methods("PUT")
	r.Handle("/commonland/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommonLand))).Methods("DELETE")
	r.Handle("/commonland/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommonLand))).Methods("POST")
}
