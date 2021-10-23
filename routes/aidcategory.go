package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AidCategoryRoutes : ""
func (route *Route) AidCategoryRoutes(r *mux.Router) {
	r.Handle("/aidcategory", Adapt(http.HandlerFunc(route.Handler.SaveAidCategory))).Methods("POST")
	r.Handle("/aidcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleAidCategory))).Methods("GET")
	r.Handle("/aidcategory", Adapt(http.HandlerFunc(route.Handler.UpdateAidCategory))).Methods("PUT")
	r.Handle("/aidcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAidCategory))).Methods("PUT")
	r.Handle("/aidcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAidCategory))).Methods("PUT")
	r.Handle("/aidcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAidCategory))).Methods("DELETE")
	r.Handle("/aidcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterAidCategory))).Methods("POST")
}
