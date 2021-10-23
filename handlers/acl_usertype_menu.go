package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//SaveACLUserTypeMenuMultiple : ""
func (h *Handler) SaveACLUserTypeMenuMultiple(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	modules := []models.ACLUserTypeMenu{}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&modules)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveACLUserTypeMenuMultiple(ctx, modules)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["update"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUserTypeMenuAccess :""
func (h *Handler) GetSingleUserTypeMenuAccess(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	userTypeID := r.URL.Query().Get("userTypeId")
	moduleID := r.URL.Query().Get("moduleId")

	if userTypeID == "" {
		response.With400V2(w, "user type id is missing", platform)
	}

	if moduleID == "" {
		response.With400V2(w, "module id is missing", platform)
	}
	module := new(models.UserTypeMenuAccess)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	module, err := h.Service.GetSingleUserTypeMenuAccess(ctx, userTypeID, moduleID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userType"] = module
	response.With200V2(w, "Success", m, platform)
}
