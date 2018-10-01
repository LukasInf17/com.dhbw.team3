package actions

import "github.com/gobuffalo/buffalo"

// InvitMailSend default implementation.
func InvitMailSend(c buffalo.Context) error {
	return c.Render(200, r.HTML("invit_mail/send.html"))
}
