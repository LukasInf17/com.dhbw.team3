package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/mailers"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// InvitMailSend sends the invitation mails to all guests.
func InvitMailSend(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	u := c.Value("current_user").(*models.User)

	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	// Allocate an empty Invitation
	invitation := &models.Invitation{}

	// To find the Invitation the parameter invitation_id is used.
	if err := tx.Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	if invitation.UserID != u.ID {
		c.Flash().Add("danger", "You are not allowed to visit this page!")
		return c.Redirect(302, "/invitations")
	}

	guests := models.Guests{}
	tx.Where("invitationid = ?", invitation.ID).All(&guests)

	if err := mailers.SendInvitMail(&guests); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("primary", "Mails were successfully sent!")
	return c.Redirect(302, "/invitations/"+c.Param("invitation_id"))
}
