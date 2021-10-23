package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SelfRegisterRoutes : ""
func (route *Route) SelfRegisterRoutes(r *mux.Router) {
	r.Handle("/selfregister", Adapt(http.HandlerFunc(route.Handler.SaveSelfRegister))).Methods("POST")
	r.Handle("/selfregister", Adapt(http.HandlerFunc(route.Handler.GetSingleSelfRegister))).Methods("GET")
	r.Handle("/selfregister", Adapt(http.HandlerFunc(route.Handler.UpdateSelfRegister))).Methods("PUT")
	r.Handle("/selfregister/filter", Adapt(http.HandlerFunc(route.Handler.FilterSelfRegister))).Methods("POST")
	r.Handle("/selfregister/status/approved", Adapt(http.HandlerFunc(route.Handler.ApprovedSelfRegister))).Methods("PUT")
	r.Handle("/selfregister/status/reject", Adapt(http.HandlerFunc(route.Handler.RejectSelfRegister))).Methods("PUT")

}
