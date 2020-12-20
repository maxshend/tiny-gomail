package main

import (
	"github.com/maxshend/tiny_gomail/logwrapper"
	pb "github.com/maxshend/tiny_gomail/tinygomail"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridSender represents a client for sending emails using Sendrid service
type SendgridSender struct {
	Key    string
	Logger *logwrapper.StandardLogger
}

// SendTextEmail sends a text email
func (s *SendgridSender) SendTextEmail(em *pb.EmailMessage) (err error) {
	err = s.sendgridEmail(em, em.Body, "")

	return
}

// SendHTMLEmail sends an HTML email
func (s *SendgridSender) SendHTMLEmail(em *pb.EmailMessage) (err error) {
	err = s.sendgridEmail(em, "", em.Body)

	return
}

func (s *SendgridSender) sendgridEmail(em *pb.EmailMessage, textContent, htmlContent string) (err error) {
	from := mail.NewEmail(em.From, em.From)
	to := mail.NewEmail(em.To[0], em.To[0])
	message := mail.NewSingleEmail(from, em.Subject, to, textContent, htmlContent)
	for i := 1; i < len(em.To); i++ {
		message.Personalizations[0].AddTos(mail.NewEmail(em.To[i], em.To[i]))
	}

	client := sendgrid.NewSendClient(s.Key)
	response, err := client.Send(message)
	s.Logger.ServiceResponse(response.Body, response.StatusCode)

	return
}
