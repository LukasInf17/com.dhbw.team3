package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Guest is the declaration of a guest
type Guest struct {
	ID                uuid.UUID `json:"id" db:"id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	InvitationID      uuid.UUID `json:"invitationid" db:"invitationid"`
	Email             string    `json:"email" db:"email"`
	Gender            int       `json:"gender" db:"gender"`
	Name              string    `json:"name" db:"name"`
	Status            int       `json:"status" db:"status"`
	AdditionalComment string    `json:"additional_comment" db:"additional_comment"`
}

type Guests []Guest
