package dao

import (
	"github.com/guregu/null"
	"zealthy-helpdesk-backend/model"
)

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

func EditUserTicket(ticketID, issueDescription string) error {
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
