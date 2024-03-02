package dao

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"zealthy-helpdesk-backend/model"
	"zealthy-helpdesk-backend/viper"
)

var DB *sqlx.DB

func DbInit() {
	DBPortString := viper.ViperReadEnvVar("DB_PORT")
	DBPort, _ := strconv.Atoi(DBPortString)

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.ViperReadEnvVar("DB_HOST"),
		DBPort,
		viper.ViperReadEnvVar("DB_USER"),
		viper.ViperReadEnvVar("DB_PASSWORD"),
		viper.ViperReadEnvVar("DB_NAME"))

	db, err := sqlx.Connect("postgres", psqlConn)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func GetAllTicketsFromUser(email string) ([]model.Ticket, error) {
	// Get all tickets from user
	SQL := `SELECT helpdesk_ticket.id, 
       helpdesk_ticket.user_id, 
       helpdesk_ticket.issue_description, 
       helpdesk_ticket.status, 
       helpdesk_ticket.admin_response, 
       helpdesk_ticket.created_at, 
       helpdesk_ticket.updated_at, 
       helpdesk_ticket.archived_at
FROM helpdesk_ticket 
JOIN users ON helpdesk_ticket.user_id = users.id
WHERE users.email = $1
`
	args := []any{email}
	var tickets []model.Ticket
	err := DB.Select(&tickets, SQL, args...)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func EditUserTicket(email, ticketID, issueDescription string) error {
	// Edit user ticket
	SQL := `UPDATE helpdesk_ticket
SET issue_description = $1
WHERE id = $2
`
	args := []any{issueDescription, ticketID}
	_, err := DB.Exec(SQL, args...)
	return err
}

func CheckUserExists(email string) (bool, error) {
	// Check if user exists
	SQL := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	args := []any{email}
	var exists bool
	err := DB.Get(&exists, SQL, args...)
	return exists, err
}

func CreateUser(email string, firstName, lastName null.String) (int64, error) {
	// Create user
	SQL := `INSERT INTO users (email, first_name, last_name) VALUES ($1, $2, $3)
INTERSECT `
	args := []any{email, firstName, lastName}
	var userID int64
	err := DB.Get(&userID, SQL, args...)
	return userID, err
}

func GetUserID(email string) (int64, error) {
	// Get user ID
	SQL := `SELECT id FROM users WHERE email = $1`
	args := []any{email}
	var userID int64
	err := DB.Get(&userID, SQL, args...)
	return userID, err
}

func CreateTicket(userID int64, issueDescription string) error {
	// Create ticket
	SQL := `INSERT INTO helpdesk_ticket (user_id, issue_description) VALUES ($1, $2)`
	args := []any{userID, issueDescription}
	_, err := DB.Exec(SQL, args...)
	return err
}

func CreateAdmin(email string, passwordHash []byte, firstName, lastName null.String) error {
	// Create admin
	SQL := `INSERT INTO admins (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)`
	args := []any{email, passwordHash, firstName, lastName}
	_, err := DB.Exec(SQL, args...)
	return err
}

func GetAllTickets() ([]model.Ticket, error) {
	// Get all tickets
	SQL := `SELECT * FROM helpdesk_ticket`
	var tickets []model.Ticket
	err := DB.Select(&tickets, SQL)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func ChangeAdminPassword(email string, passwordHash []byte) error {
	// Change admin password
	SQL := `UPDATE admins SET password = $1 WHERE email = $2`
	args := []any{string(passwordHash), email}
	_, err := DB.Exec(SQL, args...)
	return err
}

func GetAdminPasswordHash(email string) (string, error) {
	// Get admin password
	SQL := `SELECT password FROM admins WHERE email = $1`
	args := []any{email}
	var passwordHash string
	err := DB.Get(&passwordHash, SQL, args...)
	return passwordHash, err
}

func CheckAdminExists(email string) (bool, error) {
	// Check if admin exists
	SQL := `SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)`
	args := []any{email}
	var exists bool
	err := DB.Get(&exists, SQL, args...)
	return exists, err
}

func UpdateTicketStatus(ticketID int64, status string) error {
	// Update ticket status
	SQL := `UPDATE helpdesk_ticket SET status = $1 WHERE id = $2`
	args := []any{status, ticketID}
	_, err := DB.Exec(SQL, args...)
	return err
}
