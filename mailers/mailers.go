package mailers

import (
	"log"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

var smtp mail.Sender
var r *render.Engine

func init() {

	// Pulling config from the env.
	port := envy.Get("SMTP_PORT", "25")
	host := envy.Get("SMTP_HOST", "localhost")
	user := envy.Get("SMTP_USER", "NOREPLAY")
	password := envy.Get("SMTP_PASSWORD", "ybZGveyWv8x3vKvbYpy3")

	var err error
	smtp, err = mail.NewSMTPSender(host, port, user, password)

	if err != nil {
		log.Println(err)
	}

	r = render.New(render.Options{
		HTMLLayout:   "layout.html",
		TemplatesBox: packr.NewBox("../templates/mail"),
		Helpers:      render.Helpers{},
	})
}
