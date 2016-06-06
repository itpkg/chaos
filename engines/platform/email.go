package platform

import (
	"github.com/itpkg/chaos/web"
	"github.com/op/go-logging"
)

type MailSender struct {
	Logger *logging.Logger `inject:""`
}

func (p *MailSender) Send(subject, body string, html bool, files ...string) error {
	if web.IsProduction() {
		//TODO

	}
	p.Logger.Debugf("mail-to %s\n%s", subject, body)
	return nil
}
