package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BannedItemRoutes(r *mux.Router) {
	r.Handle("/bannedItem", Adapt(http.HandlerFunc(route.Handler.SaveBannedItem))).Methods("POST")
	r.Handle("/bannedItem", Adapt(http.HandlerFunc(route.Handler.GetSingleBannedItem))).Methods("GET")
	r.Handle("/bannedItem", Adapt(http.HandlerFunc(route.Handler.UpdateBannedItem))).Methods("PUT")
	r.Handle("/bannedItem/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBannedItem))).Methods("PUT")
	r.Handle("/bannedItem/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBannedItem))).Methods("PUT")
	r.Handle("/bannedItem/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBannedItem))).Methods("DELETE")
	r.Handle("/bannedItem/filter", Adapt(http.HandlerFunc(route.Handler.FilterBannedItem))).Methods("POST")
}
