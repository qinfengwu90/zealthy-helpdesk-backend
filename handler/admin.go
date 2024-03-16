package handler

import (
	"encoding/json"
	"github.com/guregu/null"
	"net/http"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/service"
	"zealthy-helpdesk-backend/utility"
)

func registerAdminHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AdminEmail string      `json:"adminEmail"`
		Password   string      `json:"password"`
		FirstName  null.String `json:"firstName"`
		LastName   null.String `json:"lastName"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreateAdmin(input.AdminEmail, input.Password, input.FirstName, input.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func changeAdminPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string `json:"email"`
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
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
		TicketID      int64       `json:"ticketId"`
		Status        null.String `json:"status"`
		AdminResponse null.String `json:"adminResponse"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = dao.UpdateTicketStatus(input.TicketID, input.Status, input.AdminResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.SendEmailUpdate(input.TicketID, input.Status, input.AdminResponse)
}

func getAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets, err := dao.GetAllTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for ticket := range tickets {
		tickets[ticket].Status = service.FormatStatus(tickets[ticket].Status)
	}
	utility.RespondJson(w, map[string]any{"tickets": tickets})
}
