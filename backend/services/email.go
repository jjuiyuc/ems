package services

import (
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/scorredoira/email"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// EmailService ...
type EmailService interface {
	SendResetEmail(c *gin.Context, name, address, token string) error
}

type defaultEmailService struct {
	cfg *viper.Viper
}

// NewEmailService ...
func NewEmailService(cfg *viper.Viper) EmailService {
	return &defaultEmailService{cfg}
}

// SendResetEmail ...
func (s defaultEmailService) SendResetEmail(c *gin.Context, name, address, token string) error {
	referrer, err := getReferrerBase(c)
	if err != nil {
		return err
	}

	body := fmt.Sprintf(`
Please follow this link to reset your password:
%s/handle-reset-link?token=%s
`, referrer, token)

	msg := email.NewMessage("Ubiik DER-EMS Password Reset", body)
	msg.AddTo(mail.Address{
		Name:    name,
		Address: address,
	})

	return sendEmail(s.cfg, msg)
}

func getReferrerBase(c *gin.Context) (referrer string, err error) {
	u, err := url.Parse(c.Request.Referer())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "getReferrerBase url.Parse",
			"err":       err,
		}).Error()
		return
	}
	referrer = u.Scheme + "://" + u.Host
	return
}

func sendEmail(cfg *viper.Viper, msg *email.Message) (err error) {
	var auth smtp.Auth
	emailServer := cfg.GetString("email.server")
	emailAccount := cfg.GetString("email.account")
	emailPassword := cfg.GetString("email.password")
	emailAccountName := cfg.GetString("email.accountName")
	emailPlainAuth := cfg.GetBool("email.plainAuth")

	server, _, err := net.SplitHostPort(emailServer)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":    "sendEmail net.SplitHostPort",
			"email.server": emailServer,
			"err":          err,
		}).Error()
		return
	}

	if emailPlainAuth {
		auth = smtp.PlainAuth("", emailAccount, emailPassword, server)
	}

	log.WithField("to", msg.Tolist()).Info("sending-email")

	msg.From = mail.Address{Name: emailAccountName, Address: emailAccount}
	err = smtp.SendMail(emailServer, auth, emailAccount, msg.Tolist(), msg.Bytes())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":     "sendEmail smtp.SendMail",
			"email.server":  emailServer,
			"email.account": emailAccount,
			"err":           err,
		}).Error("send-error")
	}
	return
}
