package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
	"zealthy-helpdesk-backend/model"
	"zealthy-helpdesk-backend/service"
	"zealthy-helpdesk-backend/utility"
)

func loginAdminHandler(w http.ResponseWriter, r *http.Request) {
	mySigningKey := jwtConfig.Secret
	var admin model.Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exists, err := service.CheckAdminExists(admin.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Admin does not exist", http.StatusUnauthorized)
		return
	}
	matches, err := service.CheckAdminPassword(admin.Email, admin.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !matches {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": admin.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.RespondJson(w, map[string]any{"token": tokenString})
}

func deleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	// Get ticket id from request body
	var input struct {
		TicketID int64 `json:"ticketId"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.DeleteTicket(input.TicketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
