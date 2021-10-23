package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) SaveAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	agroEcologicalZone := new(models.AgroEcologicalZone)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&agroEcologicalZone)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveAgroEcologicalZone(ctx, agroEcologicalZone)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["agroEcologicalZone"] = agroEcologicalZone
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) UpdateAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	agroEcologicalZone := new(models.AgroEcologicalZone)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&agroEcologicalZone)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateAgroEcologicalZone(ctx, agroEcologicalZone)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["agroEcologicalZone"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) EnableAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAgroEcologicalZone(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["agroEcologicalZone"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DisableAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAgroEcologicalZone(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["agroEcologicalZone"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DeleteAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.AgroEcologicalZone)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAgroEcologicalZone(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AgroEcologicalZone"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	agroEcologicalZone := new(models.RefAgroEcologicalZone)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	agroEcologicalZone, err := h.Service.GetSingleAgroEcologicalZone(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["agroEcologicalZone"] = agroEcologicalZone
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) FilterAgroEcologicalZone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var agroEcologicalZone *models.AgroEcologicalZoneFilter
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
	err := json.NewDecoder(r.Body).Decode(&agroEcologicalZone)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var AgroEcologicalZones []models.RefAgroEcologicalZone
	log.Println(pagination)
	AgroEcologicalZones, err = h.Service.FilterAgroEcologicalZone(ctx, agroEcologicalZone, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(AgroEcologicalZones) > 0 {
		m["agroEcologicalZone"] = AgroEcologicalZones
	} else {
		res := make([]models.AgroEcologicalZone, 0)
		m["agroEcologicalZone"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
