package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StateLiveStockRoutes : ""
func (route *Route) StateLiveStockRoutes(r *mux.Router) {
	r.Handle("/statelivestock", Adapt(http.HandlerFunc(route.Handler.SaveStateLiveStock))).Methods("POST")
	r.Handle("/statelivestock", Adapt(http.HandlerFunc(route.Handler.GetSingleStateLiveStock))).Methods("GET")
	r.Handle("/statelivestock", Adapt(http.HandlerFunc(route.Handler.UpdateStateLiveStock))).Methods("PUT")
	r.Handle("/statelivestock/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStateLiveStock))).Methods("PUT")
	r.Handle("/statelivestock/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStateLiveStock))).Methods("PUT")
	r.Handle("/statelivestock/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStateLiveStock))).Methods("DELETE")
	r.Handle("/statelivestock/filter", Adapt(http.HandlerFunc(route.Handler.FilterStateLiveStock))).Methods("POST")
}
