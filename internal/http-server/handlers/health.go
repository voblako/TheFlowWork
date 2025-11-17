package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/voblako/TheFlowWork/internal/entities"
	"github.com/voblako/TheFlowWork/storage"
)

type HealthHandler struct {
	DB      storage.DB
	Started time.Time
}

func NewHealthHandler(DB storage.DB, Started time.Time) *HealthHandler {
	return &HealthHandler{DB: DB, Started: Started}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {

	db_ver, err := h.DB.Version()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(entities.Health{
		Status:       "Ok",
		CurrentTime:  time.Now().UTC().String(),
		Uptime:       time.Since(h.Started).String(),
		Version:      "MVP",
		DatabaseInfo: db_ver,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}
