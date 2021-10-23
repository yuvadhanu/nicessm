package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MarketRoutes(r *mux.Router) {
	r.Handle("/market", Adapt(http.HandlerFunc(route.Handler.SaveMarket))).Methods("POST")
	r.Handle("/market", Adapt(http.HandlerFunc(route.Handler.GetSingleMarket))).Methods("GET")
	r.Handle("/market", Adapt(http.HandlerFunc(route.Handler.UpdateMarket))).Methods("PUT")
	r.Handle("/market/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMarket))).Methods("PUT")
	r.Handle("/market/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMarket))).Methods("PUT")
	r.Handle("/market/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMarket))).Methods("DELETE")
	r.Handle("/market/filter", Adapt(http.HandlerFunc(route.Handler.FilterMarket))).Methods("POST")
}
