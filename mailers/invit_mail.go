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
		m.From = "Invitation Factory <NOREPLY@invitation-factory.tk>"
		m.To = []string{guest.Email}
		m.SetHeader("List-Unsubscribe", "<https://invitation-factory.tk/invitations/delete_guest/"+guest.ID.String()+">")
		if err := m.AddBody(r.HTML("invit_mail.html"), render.Data{}); err != nil {
			log.Println(errors.WithStack(err))
		}
		if err := m.AddBody(r.Plain("invit_mail.txt", "layout.txt"), render.Data{}); err != nil {
			log.Println(errors.WithStack(err))
		}
		log.Println("Sending mail to " + guest.Email)
		if err := smtp.Send(m); err != nil {
			log.Println(errors.WithStack(err))
		}
	}
	return nil
}
