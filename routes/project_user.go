package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectUserRoutes : ""
func (route *Route) ProjectUserRoutes(r *mux.Router) {
	r.Handle("/projectuser", Adapt(http.HandlerFunc(route.Handler.SaveProjectUser))).Methods("POST")
	r.Handle("/projectuser", Adapt(http.HandlerFunc(route.Handler.GetSingleProjectUser))).Methods("GET")
	r.Handle("/projectuser", Adapt(http.HandlerFunc(route.Handler.UpdateProjectUser))).Methods("PUT")
	r.Handle("/projectuser/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProjectUser))).Methods("PUT")
	r.Handle("/projectuser/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProjectUser))).Methods("PUT")
	r.Handle("/projectuser/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProjectUser))).Methods("DELETE")
	r.Handle("/projectuser/filter", Adapt(http.HandlerFunc(route.Handler.FilterProjectUser))).Methods("POST")
}
