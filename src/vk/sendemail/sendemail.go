package sendemail

import (
	"net/smtp"
	"net/textproto"
	"strconv"

	vparams "vk/params"

	"github.com/jordan-wright/email"
)

func SendSMTP(subject string, text string) (err error) {
	to := vparams.Params.MessageEmailAddress
	from := `Raspberry Pi ` + vparams.Params.StationName + `<non-addrss@non-server.com>`
	host := vparams.Params.MessageSMTPHost
	user := vparams.Params.MessageSMTPUser
	pass := vparams.Params.MessageSMTPPass
	port := strconv.Itoa(vparams.Params.MessageSMTPPort)

	e := &email.Email{
		To:      []string{to},
		From:    from,
		Subject: subject,
		Text:    []byte(text),
		Headers: textproto.MIMEHeader{},
	}

	return e.Send(host+":"+port, smtp.PlainAuth("", user, pass, host))

}
