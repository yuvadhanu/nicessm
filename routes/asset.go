package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AssetRoutes : ""
func (route *Route) AssetRoutes(r *mux.Router) {
	r.Handle("/asset", Adapt(http.HandlerFunc(route.Handler.SaveAsset))).Methods("POST")
	r.Handle("/asset", Adapt(http.HandlerFunc(route.Handler.GetSingleAsset))).Methods("GET")
	r.Handle("/asset", Adapt(http.HandlerFunc(route.Handler.UpdateAsset))).Methods("PUT")
	r.Handle("/asset/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAsset))).Methods("PUT")
	r.Handle("/asset/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAsset))).Methods("PUT")
	r.Handle("/asset/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAsset))).Methods("DELETE")
	r.Handle("/asset/filter", Adapt(http.HandlerFunc(route.Handler.FilterAsset))).Methods("POST")
}
