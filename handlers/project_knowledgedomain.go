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

//SaveProjectKnowledgeDomain : ""
func (h *Handler) SaveProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	domain := new(models.ProjectKnowledgeDomain)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&domain)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveProjectKnowledgeDomain(ctx, domain)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = domain
	response.With200V2(w, "Success", m, platform)
}

//UpdateProjectKnowledgeDomain :""
func (h *Handler) UpdateProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	domain := new(models.ProjectKnowledgeDomain)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&domain)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if domain.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateProjectKnowledgeDomain(ctx, domain)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableProjectKnowledgeDomain : ""
func (h *Handler) EnableProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableProjectKnowledgeDomain(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableProjectKnowledgeDomain : ""
func (h *Handler) DisableProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableProjectKnowledgeDomain(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteProjectKnowledgeDomain : ""
func (h *Handler) DeleteProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteProjectKnowledgeDomain(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleProjectKnowledgeDomain :""
func (h *Handler) GetSingleProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ProjectKnowledgeDomain := new(models.RefProjectKnowledgeDomain)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ProjectKnowledgeDomain, err := h.Service.GetSingleProjectKnowledgeDomain(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["knowledgedomain"] = ProjectKnowledgeDomain
	response.With200V2(w, "Success", m, platform)
}

//FilterProjectKnowledgeDomain : ""
func (h *Handler) FilterProjectKnowledgeDomain(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ProjectKnowledgeDomain *models.ProjectKnowledgeDomainFilter
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
	err := json.NewDecoder(r.Body).Decode(&ProjectKnowledgeDomain)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var domains []models.RefProjectKnowledgeDomain
	log.Println(pagination)
	domains, err = h.Service.FilterProjectKnowledgeDomain(ctx, ProjectKnowledgeDomain, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(domains) > 0 {
		m["knowledgedomain"] = domains
	} else {
		res := make([]models.ProjectKnowledgeDomain, 0)
		m["knowledgedomain"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
