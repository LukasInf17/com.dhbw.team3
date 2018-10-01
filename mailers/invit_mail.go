package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendInvitMail(guest) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Invitation"
	m.From = "NORELAY@invitation-factory.tk"
	m.To = []string{}
	err := m.AddBody(r.HTML("invit_mail.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
