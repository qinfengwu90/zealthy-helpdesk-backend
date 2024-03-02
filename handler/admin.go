package handler

import (
	"encoding/json"
	"github.com/guregu/null"
	"net/http"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/service"
	"zealthy-helpdesk-backend/utility"
)

func createAdminHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserEmail string      `json:"user_email"`
		Password  string      `json:"password"`
		FirstName null.String `json:"first_name"`
		LastName  null.String `json:"last_name"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreateAdmin(input.UserEmail, input.Password, input.FirstName, input.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func changeAdminPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matches, err := service.CheckAdminPassword(input.Email, input.OldPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !matches {
		http.Error(w, "Invalid old password", http.StatusUnauthorized)
		return
	}
	err = service.ChangeAdminPassword(input.Email, input.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTicketStatusHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TicketID int64  `json:"ticket_id"`
		Status   string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = dao.UpdateTicketStatus(input.TicketID, input.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets, err := dao.GetAllTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.RespondJson(w, map[string]any{"tickets": tickets})
}
