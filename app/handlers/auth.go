package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/douglaszuqueto/controle-de-acesso/app/models"
	service "github.com/douglaszuqueto/controle-de-acesso/app/services"
	"github.com/gorilla/mux"
)

// AuthHandler auth
type AuthHandler struct{}

// AuthRequest auth
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth create endpoint
func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var request AuthRequest

	json.NewDecoder(r.Body).Decode(&request)

	res, err := service.Auth(request.Email, request.Password)

	if err != nil {
		json.NewEncoder(w).Encode(responseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(responseSuccess("", &res))
}

// Ping verify tag
func (h *AuthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tag := vars["tag"]
	idChip := vars["id_chip"]

	json.NewEncoder(w).Encode(Data{model.VerifyTag(tag, idChip)})
}
