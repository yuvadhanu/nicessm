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

//SaveCommodity : ""
func (h *Handler) SaveCommodity(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	commodity := new(models.Commodity)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = commodity
	response.With200V2(w, "Success", m, platform)
}

//UpdateCommodity :""
func (h *Handler) UpdateCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	commodity := new(models.Commodity)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if commodity.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableCommodity : ""
func (h *Handler) EnableCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableCommodity(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableCommodity : ""
func (h *Handler) DisableCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableCommodity(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteCommodity : ""
func (h *Handler) DeleteCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteCommodity(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleCommodity :""
func (h *Handler) GetSingleCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Commodity := new(models.RefCommodity)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Commodity, err := h.Service.GetSingleCommodity(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = Commodity
	response.With200V2(w, "Success", m, platform)
}

//FilterCommodity : ""
func (h *Handler) FilterCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.CommodityFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var commoditys []models.RefCommodity
	log.Println(pagination)
	commoditys, err = h.Service.FilterCommodity(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(commoditys) > 0 {
		m["commodity"] = commoditys
	} else {
		res := make([]models.Commodity, 0)
		m["commodity"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//AddInsectsCommodity :""
func (h *Handler) AddInsectsCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	commodity := new(models.Commodity)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if commodity.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.AddInsectsCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//AddDieseasesCommodity :""
func (h *Handler) AddDieseasesCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	commodity := new(models.Commodity)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if commodity.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.AddDieseasesCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteInsectsCommodity :""
func (h *Handler) DeleteInsectsCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	commodity := new(models.Commodity)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if commodity.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.DeleteInsectsCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDiseasesCommodity :""
func (h *Handler) DeleteDiseasesCommodity(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	commodity := new(models.Commodity)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&commodity)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if commodity.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.DeleteDiseasesCommodity(ctx, commodity)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["commodity"] = "success"
	response.With200V2(w, "Success", m, platform)
}
