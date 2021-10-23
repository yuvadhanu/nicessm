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

//SaveDistrictWeatherData : ""
func (h *Handler) SaveDistrictWeatherData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	districtweatherdata := new(models.DistrictWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&districtweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictWeatherData(ctx, districtweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["districtweatherdata"] = districtweatherdata
	response.With200V2(w, "Success", m, platform)
}

//UpdateDistrictWeatherData :""
func (h *Handler) UpdateDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	districtweatherdata := new(models.DistrictWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&districtweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if districtweatherdata.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDistrictWeatherData(ctx, districtweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["districtweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDistrictWeatherData : ""
func (h *Handler) EnableDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDistrictWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["districtweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDistrictWeatherData : ""
func (h *Handler) DisableDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDistrictWeatherData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["districtweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDistrictWeatherData : ""
func (h *Handler) DeleteDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.Vaccine)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDistrictWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vaccine"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDistrictWeatherData :""
func (h *Handler) GetSingleDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	districtweatherdata := new(models.RefDistrictWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	districtweatherdata, err := h.Service.GetSingleDistrictWeatherData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["districtweatherdata"] = districtweatherdata
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictWeatherData : ""
func (h *Handler) FilterDistrictWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var districtweatherdata *models.DistrictWeatherDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&districtweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var districtweatherdatas []models.RefDistrictWeatherData
	log.Println(pagination)
	districtweatherdatas, err = h.Service.FilterDistrictWeatherData(ctx, districtweatherdata, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(districtweatherdatas) > 0 {
		m["districtweatherdata"] = districtweatherdatas
	} else {
		res := make([]models.DistrictWeatherData, 0)
		m["districtweatherdata"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
