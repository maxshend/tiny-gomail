package main

import (
	"context"
	"net/smtp"
	"os"

	pb "github.com/maxshend/tiny_gomail/tinygomail"
)

type mailServer struct {
	pb.UnimplementedTinyGomailServer
}

func (m *mailServer) SendTextMessage(ctx context.Context, em *pb.EmailMessage) (*pb.SendResponse, error) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	msg := []byte("Subject: " + em.Subject + "\r\n" + em.Body + "\r\n")

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_EMAIL"), em.To, msg)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *mailServer) SendHTMLMessage(ctx context.Context, em *pb.EmailMessage) (*pb.SendResponse, error) {
	return nil, nil
}

func main() {}
