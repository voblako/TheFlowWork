package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/voblako/TheFlowWork/internal/models"
	"github.com/voblako/TheFlowWork/storage"
	"github.com/voblako/TheFlowWork/utils"
)

type CreateUserRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	ThirdName    string `json:"thirdName"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Birthdate    string `json:"birthdate"`
}

type CreateUserHandler struct {
	DB storage.DB
}

func NewCreateUserHandler(DB storage.DB) *CreateUserHandler {
	return &CreateUserHandler{DB: DB}
}

func (h *CreateUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userFromReq := CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userFromReq); err != nil {
		http.Error(w, "JSON Decode error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := utils.DateToTime(userFromReq.Birthdate)
	if err != nil {
		http.Error(w, "Not correct Birthdate format(13.04.2007): "+err.Error(), http.StatusInternalServerError)
		return
	}
	println(t.String())
	user := models.User{
		ID:           userFromReq.ID,
		Name:         userFromReq.Name,
		Surname:      userFromReq.Surname,
		ThirdName:    userFromReq.ThirdName,
		Email:        userFromReq.Email,
		PasswordHash: userFromReq.PasswordHash,
		Birthdate:    t,
	}
	err = h.DB.CreateUser(user)
	if err != nil {
		http.Error(w, "Postgres Decode error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(map[string]string{"status": "ok"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}
