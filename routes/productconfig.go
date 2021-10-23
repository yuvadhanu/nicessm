package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductConfigRoutes(r *mux.Router) {
	r.Handle("/productConfig", Adapt(http.HandlerFunc(route.Handler.SaveProductConfig))).Methods("POST")
	r.Handle("/productConfig", Adapt(http.HandlerFunc(route.Handler.GetSingleProductConfig))).Methods("GET")
	r.Handle("/productConfig", Adapt(http.HandlerFunc(route.Handler.UpdateProductConfig))).Methods("PUT")
	r.Handle("/productConfig/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProductConfig))).Methods("PUT")
	r.Handle("/productConfig/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProductConfig))).Methods("PUT")
	r.Handle("/productConfig/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProductConfig))).Methods("DELETE")
	r.Handle("/productConfig/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductConfig))).Methods("POST")
	r.Handle("/productConfig/setdefault", Adapt(http.HandlerFunc(route.Handler.SetdefaultProductConfig))).Methods("PUT")
	r.Handle("/productconfig/getdefault", Adapt(http.HandlerFunc(route.Handler.GetactiveProductConfig))).Methods("GET")
}
