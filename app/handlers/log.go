package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/douglaszuqueto/controle-de-acesso/app/models"
)

// LogHandler api endpoints
type LogHandler struct{}

// Index index endpoint
func (h *LogHandler) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess("", model.AllLogs()))
}
