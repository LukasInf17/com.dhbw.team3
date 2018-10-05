package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
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

// String is not required by pop and may be deleted
func (i Invitation) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Invitations is not required by pop and may be deleted
type Invitations []Invitation

// String is not required by pop and may be deleted
func (i Invitations) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Invitation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Invitation) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Invitation) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
