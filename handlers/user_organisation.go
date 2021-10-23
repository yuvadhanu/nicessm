package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveUserOrganisation : ""
func (h *Handler) SaveUserOrganisation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UserOrganisation := new(models.UserOrganisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&UserOrganisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserOrganisation(ctx, UserOrganisation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = UserOrganisation
	response.With200V2(w, "Success", m, platform)
}

//UpdateUserOrganisation :""
func (h *Handler) UpdateUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserOrganisation := new(models.UserOrganisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&UserOrganisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UserOrganisation.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserOrganisation(ctx, UserOrganisation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUserOrganisation : ""
func (h *Handler) EnableUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserOrganisation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableUserOrganisation : ""
func (h *Handler) DisableUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserOrganisation(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteUserOrganisation : ""
func (h *Handler) DeleteUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ID := r.URL.Query().Get("id")

	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserOrganisation(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUserOrganisation :""
func (h *Handler) GetSingleUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	UserOrganisation := new(models.RefUserOrganisation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	UserOrganisation, err := h.Service.GetSingleUserOrganisation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserOrganisation"] = UserOrganisation
	response.With200V2(w, "Success", m, platform)
}

//FilterUserOrganisation : ""
func (h *Handler) FilterUserOrganisation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var UserOrganisation *models.UserOrganisationFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&UserOrganisation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var UserOrganisations []models.RefUserOrganisation
	log.Println(pagination)
	UserOrganisations, err = h.Service.FilterUserOrganisation(ctx, UserOrganisation, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(UserOrganisations) > 0 {
		m["UserOrganisation"] = UserOrganisations
	} else {
		res := make([]models.UserOrganisation, 0)
		m["UserOrganisation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
