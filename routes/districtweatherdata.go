package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DistrictWeatherDataRoutes(r *mux.Router) {
	r.Handle("/districtWeatherData", Adapt(http.HandlerFunc(route.Handler.SaveDistrictWeatherData))).Methods("POST")
	r.Handle("/districtWeatherData", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrictWeatherData))).Methods("GET")
	r.Handle("/districtWeatherData", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherData))).Methods("PUT")
	r.Handle("/districtWeatherData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrictWeatherData))).Methods("PUT")
	r.Handle("/districtWeatherData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrictWeatherData))).Methods("PUT")
	r.Handle("/districtWeatherData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrictWeatherData))).Methods("DELETE")
	r.Handle("/districtWeatherData/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrictWeatherData))).Methods("POST")
}
