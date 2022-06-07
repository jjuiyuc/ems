package services

import (
	"fmt"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
	log "github.com/sirupsen/logrus"

	"der-ems/config"
)

// SendResetEmail ...
func SendResetEmail(name, address, referrer, token string) error {
	body := fmt.Sprintf(`
Please follow this link to reset your password:
%s/handle-reset-link?token=%s
`, referrer, token)

	msg := email.NewMessage("Ubiik DER-EMS Password Reset", body)
	msg.AddTo(mail.Address{
		Name:    name,
		Address: address,
	})

	return sendEmail(msg)
}

func sendEmail(msg *email.Message) (err error) {
	var (
		auth smtp.Auth
		cfg  = config.GetConfig()
	)

	server, _, err := net.SplitHostPort(cfg.GetString("email.server"))
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":    "sendEmail net.SplitHostPort",
			"email.server": cfg.GetString("email.server"),
			"err":          err,
		}).Error()
		return
	}

	if cfg.GetBool("email.plainAuth") {
		auth = smtp.PlainAuth("", cfg.GetString("email.account"), cfg.GetString("email.password"), server)
	}

	log.WithField("to", msg.Tolist()).Info("sending-email")

	msg.From = mail.Address{Name: cfg.GetString("email.accountName"), Address: cfg.GetString("email.account")}
	err = smtp.SendMail(cfg.GetString("email.server"), auth, cfg.GetString("email.account"), msg.Tolist(), msg.Bytes())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":     "sendEmail smtp.SendMail",
			"email.server":  cfg.GetString("email.server"),
			"email.account": cfg.GetString("email.account"),
			"err":           err,
		}).Error("send-error")
	}
	return
}
