package emailclient

import (
	cfg "ManageCenter/config"
	"fmt"

	log "github.com/inconshreveable/log15"
	"gopkg.in/gomail.v2"
)

var (
	mailUsername = cfg.Cfg.Mail.Username
	mailPassword = cfg.Cfg.Mail.Password
	mailHost     = cfg.Cfg.Mail.Host
	mailPort     = cfg.Cfg.Mail.Port
	S            gomail.SendCloser
)

func init() {

	var d = gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)
	var err error
	S, err = d.Dial()
	if err != nil {
		log.Error(fmt.Sprintf("mail err%v", err))
	}
}

func SendPlainMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]消息推广!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send plain mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendHtmlMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]消息推广!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send html mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendPackageApplyMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]用户类型变更通知!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send package apply mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}
