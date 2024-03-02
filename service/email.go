package service

import (
	"fmt"
	"zealthy-helpdesk-backend/dao"
)

func SendEmailUpdate(ticketID int64, newStatus string) error {
	emailContent := fmt.Sprintf("Your ticket #%d has been updated to status %s", ticketID, newStatus)
	return dao.StoreTicketUpdateEmail(ticketID, emailContent)
}
