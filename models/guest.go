package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Guest is the declaration of a guest
type Guest struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	InvitationID uuid.UUID `json:"invitationid" db:"invitationid"`
	Email        string    `json:"email" db:"email"`
	Gender       int       `json:"gender" db:"gender"`
	Name         string    `json:"name" db:"name"`
}

// Guests is the Guest array
type Guests []Guest
