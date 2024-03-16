package service

import (
	"fmt"
	"github.com/guregu/null"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/model"
)

func SendEmailUpdate(ticketID int64, newStatus, adminResponse null.String) error {
	emailContent := fmt.Sprintf("Your ticket #%d has been updated to ", ticketID)
	if newStatus.Valid {
		emailContent += fmt.Sprintf("status: %s", formatStatus(newStatus.String))
	}
	if adminResponse.Valid {
		emailContent += fmt.Sprintf(", with admin response: %s", adminResponse.String)
	}
	return dao.StoreTicketUpdateEmail(ticketID, emailContent)
}

func GetAllTicketsAndEmailUpdatesFromUser(email, lastName string) ([]model.Ticket, []model.Notification, error) {
	tickets, err := dao.GetAllTicketsAndFromUser(email, lastName)
	if err != nil {
		return nil, nil, err
	}
	notifications, err := dao.GetEmailUpdatesForTickets(email, lastName)
	if err != nil {
		return nil, nil, err
	}
	return tickets, notifications, nil
}

func formatStatus(status string) string {
	switch status {
	case "new":
		return "New"
	case "in_progress":
		return "In Progress"
	case "resolved":
		return "Resolved"
	default:
		return "Unknown"
	}
}
