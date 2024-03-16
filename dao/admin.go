package dao

import (
	"github.com/guregu/null"
	"zealthy-helpdesk-backend/model"
)

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
       users.email,
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
