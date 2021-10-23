package handlers

import (
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//ACLAccess :""
func (h *Handler) ACLAccess(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	userTypeID := r.URL.Query().Get("userTypeId")

	if userTypeID == "" {
		response.With400V2(w, "user type id is missing", platform)
	}
	access := new(models.ACLAccess)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	access, err := h.Service.ACLAccess(ctx, userTypeID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["access"] = access
	response.With200V2(w, "Success", m, platform)
}
