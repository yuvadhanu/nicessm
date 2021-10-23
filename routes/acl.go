package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ModuleRoutes : ""
func (route *Route) ModuleRoutes(r *mux.Router) {
	r.Handle("/acl/module", Adapt(http.HandlerFunc(route.Handler.SaveModule))).Methods("POST")
	r.Handle("/acl/module", Adapt(http.HandlerFunc(route.Handler.GetSingleModule))).Methods("GET")
	r.Handle("/acl/module", Adapt(http.HandlerFunc(route.Handler.UpdateModule))).Methods("PUT")
	r.Handle("/acl/module/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableModule))).Methods("PUT")
	r.Handle("/acl/module/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableModule))).Methods("PUT")
	r.Handle("/acl/module/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteModule))).Methods("DELETE")
	r.Handle("/acl/module/filter", Adapt(http.HandlerFunc(route.Handler.FilterModule))).Methods("POST")
}

//MenuRoutes : ""
func (route *Route) MenuRoutes(r *mux.Router) {
	r.Handle("/acl/menu", Adapt(http.HandlerFunc(route.Handler.SaveMenu))).Methods("POST")
	r.Handle("/acl/menu", Adapt(http.HandlerFunc(route.Handler.GetSingleMenu))).Methods("GET")
	r.Handle("/acl/menu", Adapt(http.HandlerFunc(route.Handler.UpdateMenu))).Methods("PUT")
	r.Handle("/acl/menu/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMenu))).Methods("PUT")
	r.Handle("/acl/menu/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMenu))).Methods("PUT")
	r.Handle("/acl/menu/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMenu))).Methods("DELETE")
	r.Handle("/acl/menu/filter", Adapt(http.HandlerFunc(route.Handler.FilterMenu))).Methods("POST")
}

//TabRoutes : ""
func (route *Route) TabRoutes(r *mux.Router) {
	r.Handle("/acl/tab", Adapt(http.HandlerFunc(route.Handler.SaveTab))).Methods("POST")
	r.Handle("/acl/tab", Adapt(http.HandlerFunc(route.Handler.GetSingleTab))).Methods("GET")
	r.Handle("/acl/tab", Adapt(http.HandlerFunc(route.Handler.UpdateTab))).Methods("PUT")
	r.Handle("/acl/tab/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTab))).Methods("PUT")
	r.Handle("/acl/tab/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTab))).Methods("PUT")
	r.Handle("/acl/tab/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTab))).Methods("DELETE")
	r.Handle("/acl/tab/filter", Adapt(http.HandlerFunc(route.Handler.FilterTab))).Methods("POST")
}

//FeatureRoutes : ""
func (route *Route) FeatureRoutes(r *mux.Router) {
	r.Handle("/acl/feature", Adapt(http.HandlerFunc(route.Handler.SaveFeature))).Methods("POST")
	r.Handle("/acl/feature", Adapt(http.HandlerFunc(route.Handler.GetSingleFeature))).Methods("GET")
	r.Handle("/acl/feature", Adapt(http.HandlerFunc(route.Handler.UpdateFeature))).Methods("PUT")
	r.Handle("/acl/feature/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFeature))).Methods("PUT")
	r.Handle("/acl/feature/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFeature))).Methods("PUT")
	r.Handle("/acl/feature/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFeature))).Methods("DELETE")
	r.Handle("/acl/feature/filter", Adapt(http.HandlerFunc(route.Handler.FilterFeature))).Methods("POST")
}

//ACLMasterUserRoutes : ""
func (route *Route) ACLMasterUserRoutes(r *mux.Router) {
	r.Handle("/acl/usertype/modules", Adapt(http.HandlerFunc(route.Handler.GetSingleModuleUserType))).Methods("GET")
	r.Handle("/acl/usertype/modules", Adapt(http.HandlerFunc(route.Handler.SaveACLUserTypeModuleMultiple))).Methods("POST")
	r.Handle("/acl/usertype/modules/filter", Adapt(http.HandlerFunc(route.Handler.FilterACLUserTypeModule))).Methods("POST")
	r.Handle("/acl/usertype/menus", Adapt(http.HandlerFunc(route.Handler.GetSingleUserTypeMenuAccess))).Methods("GET")
	r.Handle("/acl/usertype/menus", Adapt(http.HandlerFunc(route.Handler.SaveACLUserTypeMenuMultiple))).Methods("POST")

	r.Handle("/acl/usertype/tabs", Adapt(http.HandlerFunc(route.Handler.GetSingleUserTypeTabAccess))).Methods("GET")
	r.Handle("/acl/usertype/tabs", Adapt(http.HandlerFunc(route.Handler.SaveACLUserTypeTabMultiple))).Methods("POST")

	r.Handle("/acl/usertype/features", Adapt(http.HandlerFunc(route.Handler.GetSingleUserTypeFeatureAccess))).Methods("GET")
	r.Handle("/acl/usertype/features", Adapt(http.HandlerFunc(route.Handler.SaveACLUserTypeFeatureMultiple))).Methods("POST")

}

//ACLAccess : ""
func (route *Route) ACLAccess(r *mux.Router) {

	r.Handle("/acl/access", Adapt(http.HandlerFunc(route.Handler.ACLAccess))).Methods("GET")
}
