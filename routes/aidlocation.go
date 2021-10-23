package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) AidlocationRoutes(r *mux.Router) {
	r.Handle("/aidlocation", Adapt(http.HandlerFunc(route.Handler.SaveAidlocation))).Methods("POST")
	r.Handle("/aidlocation", Adapt(http.HandlerFunc(route.Handler.GetSingleAidlocation))).Methods("GET")
	r.Handle("/aidlocation", Adapt(http.HandlerFunc(route.Handler.UpdateAidlocation))).Methods("PUT")
	r.Handle("/aidlocation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAidlocation))).Methods("PUT")
	r.Handle("/aidlocation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAidlocation))).Methods("PUT")
	r.Handle("/aidlocation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAidlocation))).Methods("DELETE")
	r.Handle("/aidlocation/filter", Adapt(http.HandlerFunc(route.Handler.FilterAidlocation))).Methods("POST")
}
