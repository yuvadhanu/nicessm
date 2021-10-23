package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) InsectRoutes(r *mux.Router) {
	r.Handle("/insect", Adapt(http.HandlerFunc(route.Handler.SaveInsect))).Methods("POST")
	r.Handle("/insect", Adapt(http.HandlerFunc(route.Handler.GetSingleInsect))).Methods("GET")
	r.Handle("/insect", Adapt(http.HandlerFunc(route.Handler.UpdateInsect))).Methods("PUT")
	r.Handle("/insect/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableInsect))).Methods("PUT")
	r.Handle("/insect/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableInsect))).Methods("PUT")
	r.Handle("/insect/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteInsect))).Methods("DELETE")
	r.Handle("/insect/filter", Adapt(http.HandlerFunc(route.Handler.FilterInsect))).Methods("POST")
}
