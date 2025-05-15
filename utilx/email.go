package utilx

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/micro-services-roadmap/atom-kit/kg"
	"net/smtp"
)

// Deprecated: DoSend
// @function: send
// @description: Email发送方法
// @param: subject string, body string
// @return: error
func DoSend(to []string, subject string, body string) error {
	from := kg.C.Email.From
	nickname := kg.C.Email.Nickname
	secret := kg.C.Email.Secret
	host := kg.C.Email.Host
	port := kg.C.Email.Port
	isSSL := kg.C.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
