package service

import (
	"log"
	"net/smtp"
	"os"
	"udala/sso/exception"
)

type smtpParams struct {
	from string
	pass string
	host string
	port string
}

type SmtpService struct {
	params smtpParams
}

func NewSmtpService() *SmtpService {
	// loading env variables
	from := os.Getenv("SMTP_FROM")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	smtpParams := smtpParams{
		from: from,
		pass: pass,
		host: host,
		port: port,
	}

	return &SmtpService{
		params: smtpParams,
	}
}

func (service *SmtpService) sendEmail(email, subject, body string) exception.ApplicationException {
	auth := smtp.PlainAuth(
		"",
		service.params.from,
		service.params.pass,
		service.params.host,
	)

	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		body,
	)

	err := smtp.SendMail(
		service.params.host+":"+service.params.port,
		auth,
		service.params.from,
		to,
		msg,
	)

	if err != nil {
		log.Println("Could not send email")
		return exception.InternalServerError(err.Error())
	}

	return exception.NilError()
}
