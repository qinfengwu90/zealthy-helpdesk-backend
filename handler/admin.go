package handler

import (
	"encoding/json"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"zealthy-helpdesk-backend/dao"
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
	err = createAdmin(input.UserEmail, input.Password, input.FirstName, input.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loginAdminHandler(w http.ResponseWriter, r *http.Request) {

}

func changeAdminPasswordHandler(w http.ResponseWriter, r *http.Request) {

}

func getAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets, err := dao.GetAllTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.RespondJson(w, map[string]any{"tickets": tickets})
}

func createAdmin(email string, password string, firstName null.String, lastName null.String) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return dao.CreateAdmin(email, passwordHash, firstName, lastName)
}
