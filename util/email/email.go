package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/micro-services-roadmap/kit-common/kg"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

const (
	GmailSuffix = "@gmail.com"
	GmailType   = "google"
	GmailHost   = "smtp.gmail.com"
	GmailPort   = 587
	GmailIsSSL  = false
)

const (
	NetSuffix = "@163.com"
	NetType   = "163"
	NetHost   = "smtp.163.com"
	NetPort   = 465
	NetIsSSL  = true
)

const (
	QqSuffix = "@qq.com"
	QqType   = "qq"
	QqHost   = "smtp.qq.com"
	QqPort   = 587
	QqIsSSL  = false
)

// SendFromKgSender 发送给配置的收件人
func SendFromKgSender(subject, body, to string) error {

	return SendKgEmail(to, kg.C.Email.From, subject, body)
}

func SendToKgSender(subject, body string) error {

	return SendKgEmail(kg.C.Email.From, kg.C.Email.To, subject, body)
}

func SendKgEmail(to, from, subject, body string) error {

	return DoSend(to, subject, body, from, kg.C.Email.Nickname, kg.C.Email.Secret)
}

// DoSend 发送邮件
func DoSend(to, subject, body, from, nickname, secret string) (err error) {
	if len(from) == 0 {
		return errors.New("函数配置的发件人不能为空")
	}
	if len(to) == 0 {
		return errors.New("函数配置的收件人不能为空")
	}

	// parse host, ssl, port info
	host, port, isSSL := "", 0, false
	if strings.HasSuffix(from, GmailSuffix) {
		host = GmailHost
		port = GmailPort
		isSSL = GmailIsSSL
	} else if strings.HasSuffix(from, NetSuffix) {
		host = NetHost
		port = NetPort
		isSSL = NetIsSSL
	} else if strings.HasSuffix(from, QqSuffix) {
		host = QqHost
		port = QqPort
		isSSL = QqIsSSL
	}

	receivers := strings.Split(to, ",")
	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = receivers
	e.Subject = subject
	e.Text = []byte(body)
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
