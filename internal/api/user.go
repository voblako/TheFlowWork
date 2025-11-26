package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/voblako/TheFlowWork/internal/models"
	"github.com/voblako/TheFlowWork/storage"
	"github.com/voblako/TheFlowWork/utils"
)

type UserRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	ThirdName    string `json:"thirdName"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Birthdate    string `json:"birthdate"`
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	userFromReq := UserRequest{}
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
	err = s.DB.CreateUser(user)
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
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.ParseInt(r.PathValue("user_id"), 10, 32)
	if err != nil {
		http.Error(w, "User`s id is not correct: "+err.Error(), http.StatusBadRequest)
	}

	user, err := s.DB.GetUser(int32(user_id))
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResp := UserRequest{
		ID:           user.ID,
		Name:         user.Name,
		Surname:      user.Surname,
		ThirdName:    user.ThirdName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Birthdate:    utils.TimeToDate(user.Birthdate),
	}

	jsonResp, err := json.Marshal(userResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
