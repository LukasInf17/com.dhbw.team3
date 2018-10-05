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

// SendVerifyMail sends an email where the User verifies his email address
func SendVerifyMail(u *models.User) error {
	m := mail.NewMessage()
	m.Subject = "Verify your email address to gain access to the Invitation Factory"
	m.From = "Invitation Factory <NOREPLY@invitation-factory.tk>"
	m.To = []string{u.Email}
	m.SetHeader("List-Unsubscribe", "<https://invitation-factory.tk/invitations/delete_guest/"+u.ID.String()+">")
	textbody := `Welcome to the Invitation Factory!

It seems that someone with the email address ` + u.Email + ` tried to register to the Invitation Factory.
When it was you, you can now verify your email address here: https://invitation-factory.tk/users/` + u.ID.String() + `/verify


Yours sincerely,

The Invitation Factory team`

	if err := m.AddBody(r.HTML("verify_mail.html"), render.Data{"userid": u.ID}); err != nil {
		log.Println(errors.WithStack(err))
	}
	if err := m.AddBody(r.String(textbody), render.Data{}); err != nil {
		log.Println(errors.WithStack(err))
	}
	log.Println("Sending verify mail to " + u.Email)
	if err := smtp.Send(m); err != nil {
		log.Println(errors.WithStack(err))
	}
	return nil
}
