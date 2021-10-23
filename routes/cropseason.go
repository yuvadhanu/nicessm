package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CropseasonRoutes(r *mux.Router) {
	r.Handle("/cropseason", Adapt(http.HandlerFunc(route.Handler.SaveCropseason))).Methods("POST")
	r.Handle("/cropseason", Adapt(http.HandlerFunc(route.Handler.GetSingleCropseason))).Methods("GET")
	r.Handle("/cropseason", Adapt(http.HandlerFunc(route.Handler.UpdateCropseason))).Methods("PUT")
	r.Handle("/cropseason/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCropseason))).Methods("PUT")
	r.Handle("/cropseason/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCropseason))).Methods("PUT")
	r.Handle("/cropseason/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCropseason))).Methods("DELETE")
	r.Handle("/cropseason/filter", Adapt(http.HandlerFunc(route.Handler.FilterCropseason))).Methods("POST")
}
