package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProjectKnowledgeDomainRoutes : ""
func (route *Route) ProjectKnowledgeDomainRoutes(r *mux.Router) {
	r.Handle("/projectknowledgedomain", Adapt(http.HandlerFunc(route.Handler.SaveProjectKnowledgeDomain))).Methods("POST")
	r.Handle("/projectknowledgedomain", Adapt(http.HandlerFunc(route.Handler.GetSingleProjectKnowledgeDomain))).Methods("GET")
	r.Handle("/projectknowledgedomain", Adapt(http.HandlerFunc(route.Handler.UpdateProjectKnowledgeDomain))).Methods("PUT")
	r.Handle("/projectknowledgedomain/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProjectKnowledgeDomain))).Methods("PUT")
	r.Handle("/projectknowledgedomain/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProjectKnowledgeDomain))).Methods("PUT")
	r.Handle("/projectknowledgedomain/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProjectKnowledgeDomain))).Methods("DELETE")
	r.Handle("/projectknowledgedomain/filter", Adapt(http.HandlerFunc(route.Handler.FilterProjectKnowledgeDomain))).Methods("POST")
}

//KnowledgeDomainRoutes : ""
func (route *Route) KnowledgeDomainRoutes(r *mux.Router) {
	r.Handle("/knowledgedomain", Adapt(http.HandlerFunc(route.Handler.SaveKnowledgeDomain))).Methods("POST")
	r.Handle("/knowledgedomain", Adapt(http.HandlerFunc(route.Handler.GetSingleKnowledgeDomain))).Methods("GET")
	r.Handle("/knowledgedomain", Adapt(http.HandlerFunc(route.Handler.UpdateKnowledgeDomain))).Methods("PUT")
	r.Handle("/knowledgedomain/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableKnowledgeDomain))).Methods("PUT")
	r.Handle("/knowledgedomain/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableKnowledgeDomain))).Methods("PUT")
	r.Handle("/knowledgedomain/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteKnowledgeDomain))).Methods("DELETE")
	r.Handle("/knowledgedomain/filter", Adapt(http.HandlerFunc(route.Handler.FilterKnowledgeDomain))).Methods("POST")
}

//subDomainRoutes : ""
func (route *Route) SubDomainRoutes(r *mux.Router) {
	r.Handle("/subdomain", Adapt(http.HandlerFunc(route.Handler.SaveSubDomain))).Methods("POST")
	r.Handle("/subdomain", Adapt(http.HandlerFunc(route.Handler.GetSingleSubDomain))).Methods("GET")
	r.Handle("/subdomain", Adapt(http.HandlerFunc(route.Handler.UpdateSubDomain))).Methods("PUT")
	r.Handle("/subdomain/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSubDomain))).Methods("PUT")
	r.Handle("/subdomain/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSubDomain))).Methods("PUT")
	r.Handle("/subdomain/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSubDomain))).Methods("DELETE")
	r.Handle("/subdomain/filter", Adapt(http.HandlerFunc(route.Handler.FilterSubDomain))).Methods("POST")
}

//TopicRoutes : ""
func (route *Route) TopicRoutes(r *mux.Router) {
	r.Handle("/topic", Adapt(http.HandlerFunc(route.Handler.SaveTopic))).Methods("POST")
	r.Handle("/topic", Adapt(http.HandlerFunc(route.Handler.GetSingleTopic))).Methods("GET")
	r.Handle("/topic", Adapt(http.HandlerFunc(route.Handler.UpdateTopic))).Methods("PUT")
	r.Handle("/topic/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTopic))).Methods("PUT")
	r.Handle("/topic/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTopic))).Methods("PUT")
	r.Handle("/topic/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTopic))).Methods("DELETE")
	r.Handle("/topic/filter", Adapt(http.HandlerFunc(route.Handler.FilterTopic))).Methods("POST")
}

//SubTopicRoutes : ""
func (route *Route) SubTopicRoutes(r *mux.Router) {
	r.Handle("/subtopic", Adapt(http.HandlerFunc(route.Handler.SaveSubTopic))).Methods("POST")
	r.Handle("/subtopic", Adapt(http.HandlerFunc(route.Handler.GetSingleSubTopic))).Methods("GET")
	r.Handle("/subtopic", Adapt(http.HandlerFunc(route.Handler.UpdateSubTopic))).Methods("PUT")
	r.Handle("/subtopic/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSubTopic))).Methods("PUT")
	r.Handle("/subtopic/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSubTopic))).Methods("PUT")
	r.Handle("/subtopic/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSubTopic))).Methods("DELETE")
	r.Handle("/subtopic/filter", Adapt(http.HandlerFunc(route.Handler.FilterSubTopic))).Methods("POST")
}
