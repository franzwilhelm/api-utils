package utils

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridAPIKey stores the token to ensure a valid connection to the sendgrid API
var SendgridAPIKey string

// SendMail is used to send a mail  through the sendgrid API
func SendMail(subject, message, toMail string) (*rest.Response, error) {
	from := mail.NewEmail("Stempl.no", "noreply@stempl.no")
	to := mail.NewEmail("Franz", toMail)

	content := mail.NewContent("text/plain", message)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(SendgridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	return sendgrid.API(request)
}

// ValidateHost is used to check if an email is valid by checking MX records ++
func ValidateHost(host string) (err error) {
	fmt.Println(host)
	err = checkmail.ValidateFormat(host)
	if err != nil {
		return fmt.Errorf("Email error: %v", err)
	}
	err = checkmail.ValidateHost(host)
	if err != nil {
		return fmt.Errorf("Email error: %v", err)
	}
	if errS, _ := err.(checkmail.SmtpError); errS.Err != nil {
		err = fmt.Errorf("Email error: %v", errS)
	}
	return
}
