package handlers

import (
	"encoding/json"
	"net/http"

	model "../models"
)

// LogHandler api endpoints
type LogHandler struct{}

// Index index endpoint
func (h *LogHandler) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responseSuccess("", model.AllLogs()))
}
