package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	model "github.com/douglaszuqueto/controle-de-acesso/app/models"

	"github.com/gorilla/mux"
)

// DeviceHandler api endpoints
type DeviceHandler struct{}

// DeviceTag create/update
type DeviceTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ChipID      string `json:"chip_id"`
	State       int    `json:"state"`
}

// Index index endpoint
func (h *DeviceHandler) Index(w http.ResponseWriter, r *http.Request) {
	devices := model.AllDevices()

	json.NewEncoder(w).Encode(responseSuccess("", devices))
}

// Get get endpoint
func (h *DeviceHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	json.NewEncoder(w).Encode(responseSuccess("", model.FindDevice(id)))
}

// Tags tags endpoint
func (h *DeviceHandler) Tags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	json.NewEncoder(w).Encode(responseSuccess("", model.FindTags(id)))
}

// Create create endpoint
func (h *DeviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request DeviceTag

	json.NewDecoder(r.Body).Decode(&request)

	var chipID sql.NullString
	chipID.String = request.ChipID

	device := &model.Device{
		Name:        request.Name,
		Description: request.Description,
		ChipID:      chipID,
		State:       request.State}

	device, err := model.CreateDevice(*device)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &device))
}

// Update update endpoint
func (h *DeviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request DeviceTag
	vars := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&request)

	var chipID sql.NullString
	chipID.String = request.ChipID

	device := &model.Device{
		Name:        request.Name,
		Description: request.Description,
		ChipID:      chipID,
		State:       request.State}

	id := vars["id"]
	device, err := model.UpdateDevice(id, *device)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &device))
}

// Remove remove endpoint
func (h *DeviceHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := model.RemoveDevice(id)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("Device removido com sucesso", nil))
}

// TagActivateDeactivate tagActivateDeactivate endpoint
func (h *DeviceHandler) TagActivateDeactivate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idTag := vars["id_tag"]

	err := model.TagActivateDeactivate(id, idTag)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("Acesso atualizado com sucesso", nil))
}
