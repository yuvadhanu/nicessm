package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OrganisationRoutes : ""
func (route *Route) OrganisationRoutes(r *mux.Router) {
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.SaveOrganisation))).Methods("POST")
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.GetSingleOrganisation))).Methods("GET")
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.UpdateOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrganisation))).Methods("DELETE")
	r.Handle("/organisation/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrganisation))).Methods("POST")
}

//UserOrganisationRoutes : ""
func (route *Route) UserOrganisationRoutes(r *mux.Router) {
	r.Handle("/userorganisation", Adapt(http.HandlerFunc(route.Handler.SaveUserOrganisation))).Methods("POST")
	r.Handle("/userorganisation", Adapt(http.HandlerFunc(route.Handler.GetSingleUserOrganisation))).Methods("GET")
	r.Handle("/userorganisation", Adapt(http.HandlerFunc(route.Handler.UpdateUserOrganisation))).Methods("PUT")
	r.Handle("/userorganisation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserOrganisation))).Methods("PUT")
	r.Handle("/userorganisation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserOrganisation))).Methods("PUT")
	r.Handle("/userorganisation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserOrganisation))).Methods("DELETE")
	r.Handle("/userorganisation/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserOrganisation))).Methods("POST")
}

//UserRoutes : ""
func (route *Route) UserRoutes(r *mux.Router) {
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.SaveUser))).Methods("POST")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.GetSingleUser))).Methods("GET")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.UpdateUser))).Methods("PUT")
	r.Handle("/user/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUser))).Methods("PUT")
	r.Handle("/user/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUser))).Methods("PUT")
	r.Handle("/user/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUser))).Methods("DELETE")
	r.Handle("/user/filter", Adapt(http.HandlerFunc(route.Handler.FilterUser))).Methods("POST")
	r.Handle("/user/resetpassword", Adapt(http.HandlerFunc(route.Handler.ResetUserPassword))).Methods("PUT")
	r.Handle("/user/changepassword", Adapt(http.HandlerFunc(route.Handler.ChangePassword))).Methods("PUT")
	r.Handle("/user/forgetpassword/generateotp", Adapt(http.HandlerFunc(route.Handler.ForgetPasswordGenerateOTP))).Methods("GET")
	r.Handle("/user/forgetpassword/validateotp", Adapt(http.HandlerFunc(route.Handler.ForgetPasswordValidateOTP))).Methods("GET")
	r.Handle("/user/passwordupdate", Adapt(http.HandlerFunc(route.Handler.PasswordUpdate))).Methods("PUT")
	//UserCollectionLimit updates
	r.Handle("/user/collectionlimitupdate", Adapt(http.HandlerFunc(route.Handler.UserCollectionLimit))).Methods("PUT")

}

//UserAuthRoutes : ""
func (route *Route) UserAuthRoutes(r *mux.Router) {
	r.Handle("/user/auth", Adapt(http.HandlerFunc(route.Handler.Login))).Methods("POST")
	r.Handle("/user/auth/otplogin/generateotp", Adapt(http.HandlerFunc(route.Handler.OTPLoginGenerateOTP))).Methods("POST")
	r.Handle("/user/auth/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.OTPLoginValidateOTP))).Methods("POST")
}

//UserTypeRoutes : ""
func (route *Route) UserTypeRoutes(r *mux.Router) {
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.SaveUserType))).Methods("POST")
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.GetSingleUserType))).Methods("GET")
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.UpdateUserType))).Methods("PUT")
	r.Handle("/usertype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserType))).Methods("PUT")
	r.Handle("/usertype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserType))).Methods("PUT")
	r.Handle("/usertype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserType))).Methods("DELETE")
	r.Handle("/usertype/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserType))).Methods("POST")
}
