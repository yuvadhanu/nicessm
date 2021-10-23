package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//BlockCropRoutes : ""
func (route *Route) BlockCropRoutes(r *mux.Router) {
	r.Handle("/blockcrop", Adapt(http.HandlerFunc(route.Handler.SaveBlockCrop))).Methods("POST")
	r.Handle("/blockcrop", Adapt(http.HandlerFunc(route.Handler.GetSingleBlockCrop))).Methods("GET")
	r.Handle("/blockcrop", Adapt(http.HandlerFunc(route.Handler.UpdateBlockCrop))).Methods("PUT")
	r.Handle("/blockcrop/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBlockCrop))).Methods("PUT")
	r.Handle("/blockcrop/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBlockCrop))).Methods("PUT")
	r.Handle("/blockcrop/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBlockCrop))).Methods("DELETE")
	r.Handle("/blockcrop/filter", Adapt(http.HandlerFunc(route.Handler.FilterBlockCrop))).Methods("POST")
}
