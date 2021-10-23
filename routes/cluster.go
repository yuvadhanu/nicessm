package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ClusterRoutes : ""
func (route *Route) ClusterRoutes(r *mux.Router) {
	r.Handle("/cluster", Adapt(http.HandlerFunc(route.Handler.SaveCluster))).Methods("POST")
	r.Handle("/cluster", Adapt(http.HandlerFunc(route.Handler.GetSingleCluster))).Methods("GET")
	r.Handle("/cluster", Adapt(http.HandlerFunc(route.Handler.UpdateCluster))).Methods("PUT")
	r.Handle("/cluster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCluster))).Methods("PUT")
	r.Handle("/cluster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCluster))).Methods("PUT")
	r.Handle("/cluster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCluster))).Methods("DELETE")
	r.Handle("/cluster/filter", Adapt(http.HandlerFunc(route.Handler.FilterCluster))).Methods("POST")
}
