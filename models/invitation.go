package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Invitation is the structure
type Invitation struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	UserID       uuid.UUID `json:"userid" db:"userid"`
	Salutation   int       `json:"salutation" db:"salutation"`
	Mailtext     string    `json:"mailtext" db:"mailtext"`
	SentToGuests bool      `json:"senttoguests" db:"senttoguests"`
	Guests       Guests    `has_many:"guests" fk_id:"invitationid"`
}

// Invitations is not required by pop and may be deleted
type Invitations []Invitation
