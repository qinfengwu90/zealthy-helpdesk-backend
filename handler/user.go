package handler

import (
	"encoding/json"
	"github.com/guregu/null"
	"net/http"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/service"
	"zealthy-helpdesk-backend/utility"
)

func getAllTicketsAndEmailUpdatesForUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user email from request body
	var input struct {
		UserEmail string `json:"userEmail"`
		LastName  string `json:"lastName"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get all tickets from user
	tickets, email, err := service.GetAllTicketsAndEmailUpdatesFromUser(input.UserEmail, input.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return all tickets
	utility.RespondJson(w, map[string]any{"tickets": tickets, "emails": email})
}

func editUserTicketHandler(w http.ResponseWriter, r *http.Request) {
	// Get ticket info from request body
	var input struct {
		TicketID         string `json:"ticketId"`
		IssueDescription string `json:"issueDescription"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dao.EditUserTicket(input.TicketID, input.IssueDescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createTicketHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserEmail        string      `json:"userEmail"`
		IssueDescription string      `json:"issueDescription"`
		FirstName        null.String `json:"firstName"`
		LastName         null.String `json:"lastName"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = createTicket(input.UserEmail, input.IssueDescription, input.FirstName, input.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTicket(userEmail, issueDescription string, firstName, lastName null.String) error {
	// Check if user exists
	userExists, err := dao.CheckUserExists(userEmail)
	if err != nil {
		return err
	}

	// If not, create user
	var userID int64
	if !userExists {
		userID, err = dao.CreateUser(userEmail, firstName, lastName)
		if err != nil {
			return err
		}
	} else {
		// Get user ID
		userID, err = dao.GetUserID(userEmail)
		if err != nil {
			return err
		}
	}

	// Create ticket
	err = dao.CreateTicket(userID, issueDescription)
	if err != nil {
		return err
	}
	return nil
}
