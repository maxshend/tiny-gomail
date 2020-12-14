package main

import (
	"context"
	"log"
	"net"
	"net/smtp"
	"os"
	"sync"

	pb "github.com/maxshend/tiny_gomail/tinygomail"
	"google.golang.org/grpc"
)

const htmlMIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
const defaultHost = "localhost"
const defaultPort = "8000"

type mailServer struct {
	pb.UnimplementedTinyGomailServer
	mu sync.Mutex
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

func main() {
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = defaultPort
	}

	host, exists := os.LookupEnv("SERVER_HOST")
	if !exists {
		host = defaultHost
	}

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterTinyGomailServer(grpcServer, &mailServer{})
	grpcServer.Serve(lis)
}
