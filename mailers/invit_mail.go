package mailers

import (
	"log"

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
		if err := m.AddBody(r.HTML("invit_mail.html"), render.Data{}); err != nil {
			log.Println(errors.WithStack(err))
		}
		log.Println("Sending mail to " + guest.Email)
		if err := smtp.Send(m); err != nil {
			log.Println(errors.WithStack(err))
		}
	}
	return nil
}
