package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// SendInvitMail sends a mail to all guests of the invitation
func SendInvitMail(guests *models.Guests) error {
	for _, guest := range *guests {
		m := mail.NewMessage()
		m.Subject = "Invitation"
		m.From = "NOREPLAY@invitation-factory.tk"
		m.To = []string{guest.Email}
		err := m.AddBody(r.HTML("invit_mail.html"), render.Data{})
		if err != nil {
			return errors.WithStack(err)
		}
		return smtp.Send(m)
	}
	return nil
}
