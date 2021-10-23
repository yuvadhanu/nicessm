package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DiseaseRoutes(r *mux.Router) {
	r.Handle("/disease", Adapt(http.HandlerFunc(route.Handler.SaveDisease))).Methods("POST")
	r.Handle("/disease", Adapt(http.HandlerFunc(route.Handler.GetSingleDisease))).Methods("GET")
	r.Handle("/disease", Adapt(http.HandlerFunc(route.Handler.UpdateDisease))).Methods("PUT")
	r.Handle("/disease/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDisease))).Methods("PUT")
	r.Handle("/disease/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDisease))).Methods("PUT")
	r.Handle("/disease/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDisease))).Methods("DELETE")
	r.Handle("/disease/filter", Adapt(http.HandlerFunc(route.Handler.FilterDisease))).Methods("POST")
}
