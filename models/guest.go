package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

type Guest struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	InvitationID uuid.UUID `json:"invitationid" db:"invitationid"`
	Email        int       `json:"email" db:"email"`
	Gender       int8      `json:"gender" db:"gender"`
	Name         string    `json:"name" db:"name"`
}
