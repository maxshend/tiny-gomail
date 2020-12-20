package main

import (
	"net/smtp"

	"github.com/maxshend/tiny_gomail/logwrapper"
	pb "github.com/maxshend/tiny_gomail/tinygomail"
)

// SMTPSender represents a client for sending emails using SMTP
type SMTPSender struct {
	Email    string
	Password string
	Host     string
	Port     string
	Logger   *logwrapper.StandardLogger
}

// SendTextEmail sends a text email
func (s *SMTPSender) SendTextEmail(em *pb.EmailMessage) (err error) {
	err = s.sendSMTPEmail(em, "")

	return
}

// SendHTMLEmail sends an HTML email
func (s *SMTPSender) SendHTMLEmail(em *pb.EmailMessage) (err error) {
	err = s.sendSMTPEmail(em, htmlMIME)

	return
}

func (s *SMTPSender) sendSMTPEmail(em *pb.EmailMessage, mime string) (err error) {
	msg := []byte("Subject: " + em.Subject + "\r\n" + mime + "\r\n" + em.Body + "\r\n")
	auth := smtp.PlainAuth("", s.Email, s.Password, s.Host)
	err = smtp.SendMail(s.Host+":"+s.Port, auth, s.Email, em.To, msg)

	return
}
