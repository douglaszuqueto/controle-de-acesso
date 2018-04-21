package handlers

import (
	"encoding/json"
	"net/http"

	model "../models"
	"github.com/gorilla/mux"
)

// TagHandler api endpoints
type TagHandler struct{}

// TagRequest create/update
type TagRequest struct {
	IDUser      string `json:"id_user"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	State       int    `json:"state"`
}

// Index index endpoint
func (h *TagHandler) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess("", model.AllTags()))
}

// Get get endpoint
func (h *TagHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	t, err := model.FindTag(id)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", t))
}

// Create create endpoint
func (h *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request TagRequest

	json.NewDecoder(r.Body).Decode(&request)

	tag := &model.Tag{
		IDUser:      request.IDUser,
		Tag:         request.Tag,
		Description: request.Description,
		State:       request.State}

	tag, err := model.CreateTag(*tag)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &tag))
}

// Update update endpoint
func (h *TagHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request TagRequest
	vars := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&request)

	tag := &model.Tag{
		IDUser:      request.IDUser,
		Tag:         request.Tag,
		Description: request.Description,
		State:       request.State}

	id := vars["id"]
	tag, err := model.UpdateTag(id, *tag)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &tag))
}

// Remove remove endpoint
func (h *TagHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := model.RemoveTag(id)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("Removido com sucesso", err))
}

// TagDeviceRequest create/update
type TagDeviceRequest struct {
	IDDevice string `json:"id_device"`
	IDTag    string `json:"id_tag"`
}

// AttachTagDevice attaching tag device
func (h *TagHandler) AttachTagDevice(w http.ResponseWriter, r *http.Request) {
	var request TagDeviceRequest

	json.NewDecoder(r.Body).Decode(&request)

	tagDevice := &model.TagDevice{
		IDDevice: request.IDDevice,
		IDTag:    request.IDTag}

	res, err := model.AttachTagDevice(*tagDevice)

	if !res {
		json.NewEncoder(w).Encode(responseError("Algum erro aconteceu!"))
		return
	}

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("Tag associada com sucesso!", nil))
}

// DettachTagDevice dettaching tag device
func (h *TagHandler) DettachTagDevice(w http.ResponseWriter, r *http.Request) {
	var request TagDeviceRequest

	json.NewDecoder(r.Body).Decode(&request)

	tagDevice := &model.TagDevice{
		IDDevice: request.IDDevice,
		IDTag:    request.IDTag}

	err := model.DettachTagDevice(*tagDevice)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("Tag desassociada com sucesso!", nil))
}
