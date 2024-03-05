package model

import (
	"github.com/guregu/null"
	"time"
)

type User struct {
	Id         int64       `json:"id"`
	Email      string      `json:"email"`
	FirstName  null.String `json:"first_name"`
	LastName   null.String `json:"last_name"`
	CreatedAt  time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time   `json:"updatedAt" db:"updated_at"`
	ArchivedAt null.Time   `json:"archivedAt" db:"archived_at"`
}

type Admin struct {
	Id         int64       `json:"id"`
	Email      string      `json:"email"`
	Salt       string      `json:"salt"`
	Password   string      `json:"password"`
	FirstName  null.String `json:"firstName"`
	LastName   null.String `json:"lastName"`
	CreatedAt  time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time   `json:"updatedAt" db:"updated_at"`
	ArchivedAt null.Time   `json:"archivedAt" db:"archived_at"`
}

type Ticket struct {
	Id               int64       `json:"id" db:"id"`
	UserId           int64       `json:"userId" db:"user_id"`
	FirstName        string      `json:"firstName" db:"first_name"`
	LastName         string      `json:"lastName" db:"last_name"`
	IssueDescription string      `json:"issueDescription" db:"issue_description"`
	Status           string      `json:"status" db:"status"`
	AdminResponse    null.String `json:"adminResponse" db:"admin_response"`
	CreatedAt        time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time   `json:"updatedAt" db:"updated_at"`
	ArchivedAt       null.Time   `json:"-" db:"archived_at"`
}

type Notification struct {
	Id        int64     `json:"id" db:"id"`
	TicketId  int64     `json:"ticketId" db:"ticket_id"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
