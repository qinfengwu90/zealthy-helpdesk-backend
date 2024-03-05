package dao

import (
	"fmt"
	"github.com/guregu/null"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"zealthy-helpdesk-backend/model"
	"zealthy-helpdesk-backend/utility"
)

var DB *sqlx.DB

func DbInit(dbConfig *utility.PostgresInfo) {
	DBPortString := dbConfig.Port
	DBPort, _ := strconv.Atoi(DBPortString)

	unixSocketPath := fmt.Sprintf("/cloudsql/%s/.s.PGSQL.5432", dbConfig.CloudSqlConnectionName)

	psqSocketConn := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true",
		dbConfig.Username,
		dbConfig.Password,
		unixSocketPath,
		dbConfig.Dbname)

	db, err := sqlx.Connect("postgres", psqSocketConn)

	if err != nil {
		psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			DBPort,
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Dbname)

		db, err := sqlx.Connect("postgres", psqlConn)
		if err != nil {
			log.Fatal(err)
		}
		DB = db

	} else {
		DB = db
	}
}

func GetAllTicketsAndFromUser(email, lastName string) ([]model.Ticket, error) {
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
  AND users.last_name = $2
  AND helpdesk_ticket.archived_at IS NULL
ORDER BY helpdesk_ticket.updated_at DESC
`
	args := []any{email, lastName}
	var tickets []model.Ticket
	err := DB.Select(&tickets, SQL, args...)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetEmailUpdatesForTickets(email, lastName string) ([]model.Notification, error) {
	// Get email updates for tickets
	SQL := `SELECT ticket_notification_email.id, 
	   ticket_notification_email.ticket_id, 
	   ticket_notification_email.message, 
	   ticket_notification_email.created_at
FROM ticket_notification_email
JOIN helpdesk_ticket ON ticket_notification_email.ticket_id = helpdesk_ticket.id
JOIN users ON helpdesk_ticket.user_id = users.id
WHERE users.email = $1
  AND users.last_name = $2
ORDER BY ticket_notification_email.created_at DESC
`
	args := []any{email, lastName}
	var notifications []model.Notification
	err := DB.Select(&notifications, SQL, args...)
	if err != nil {
		return nil, err
	}
	return notifications, nil
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
RETURNING id`
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
	SQL := `SELECT helpdesk_ticket.id,
       helpdesk_ticket.user_id,
       users.first_name,
       users.last_name,
       helpdesk_ticket.issue_description,
       helpdesk_ticket.status,
       helpdesk_ticket.admin_response,
       helpdesk_ticket.created_at,
       helpdesk_ticket.updated_at
FROM helpdesk_ticket
JOIN users ON helpdesk_ticket.user_id = users.id
WHERE helpdesk_ticket.archived_at IS NULL
ORDER BY helpdesk_ticket.updated_at
`
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

func UpdateTicketStatus(ticketID int64, status, adminResponse null.String) error {
	// Update ticket status
	SQL := `UPDATE helpdesk_ticket 
SET status = $1, admin_response = $2
WHERE id = $3`
	args := []any{status, adminResponse, ticketID}
	_, err := DB.Exec(SQL, args...)
	return err
}

func StoreTicketUpdateEmail(ticketID int64, message string) error {
	// Store ticket update email
	SQL := `INSERT INTO ticket_notification_email (ticket_id, message) VALUES ($1, $2)`
	args := []any{ticketID, message}
	_, err := DB.Exec(SQL, args...)
	return err
}

func DeleteTicket(ticketID int64) error {
	// Delete ticket
	SQL := `UPDATE helpdesk_ticket SET archived_at = NOW() WHERE id = $1`
	args := []any{ticketID}
	_, err := DB.Exec(SQL, args...)
	return err
}
