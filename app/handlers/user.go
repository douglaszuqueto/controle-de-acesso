package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	model "../models"
)

// UserHandler api endpoints
type UserHandler struct{}

// Index index endpoint
func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess("", model.AllUsers()))
}

// Get get endpoint
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	m, err := model.FindUser(id)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", m))
}

// UserRequest create/update
type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create create endpoint
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request UserRequest

	json.NewDecoder(r.Body).Decode(&request)

	user := &model.User{
		Name:  request.Name,
		Email: request.Email}

	user, err := model.CreateUser(*user)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &user))
}

// Update update endpoint
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request UserRequest
	vars := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&request)

	user := &model.User{
		Name:  request.Name,
		Email: request.Email}

	id := vars["id"]

	user, err := model.UpdateUser(id, *user)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(responseSuccess("", &user))
}

// Remove remove endpoint
func (h *UserHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := model.RemoveUser(id)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(responseSuccess("Usuário removido com sucesso", nil))
}

// GetUserTags get endpoint
func (h *UserHandler) GetUserTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	json.NewEncoder(w).Encode(responseSuccess("", model.GetUserTags(id)))
}
