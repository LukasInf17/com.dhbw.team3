package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// DeleteGuestFromUnsubscribe deletes a guest when he unsubscribes
func DeleteGuestFromUnsubscribe(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	guest := &models.Guest{}
	guests := models.Guests{}

	if err := tx.Find(guest, c.Param("guest_id")); err != nil {
		return c.Error(404, err)
	}

	tx.Where("email = ?", guest.Email).All(&guests)

	if err := tx.Destroy(&guests); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.String("Your e-mail was deleted successfully"))
}

// StatusResponse serves the page where the guest can response to the invitation.
func StatusResponse(c buffalo.Context) error {
	return c.Render(200, r.HTML("guests/response.html"))
}

// SetStatusResponse puts the guest response data into the database.
func SetStatusResponse(c buffalo.Context) error {

	return c.Redirect(302, "/")
}
