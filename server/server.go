package main

import (
	"context"
	"net/smtp"
	"os"

	pb "github.com/maxshend/tiny_gomail/tinygomail"
)

const htmlMIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"

type mailServer struct {
	pb.UnimplementedTinyGomailServer
}

func (m *mailServer) SendTextMessage(ctx context.Context, em *pb.EmailMessage) (response *pb.SendResponse, err error) {
	err = sendSMTPEmail(em, "")

	return
}

func (m *mailServer) SendHTMLMessage(ctx context.Context, em *pb.EmailMessage) (response *pb.SendResponse, err error) {
	err = sendSMTPEmail(em, htmlMIME)

	return
}

func sendSMTPEmail(em *pb.EmailMessage, mime string) (err error) {
	msg := []byte("Subject: " + em.Subject + "\r\n" + mime + "\r\n" + em.Body + "\r\n")
	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_EMAIL"), em.To, msg)

	return
}

func main() {}
