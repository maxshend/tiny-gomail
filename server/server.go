package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	"github.com/maxshend/tiny_gomail/logwrapper"
	pb "github.com/maxshend/tiny_gomail/tinygomail"
	"google.golang.org/grpc"
)

const defaultHost = "localhost"
const defaultPort = "8000"

type mailServer struct {
	pb.UnimplementedTinyGomailServer
	mu     sync.Mutex
	sender Sender
}

func (m *mailServer) SendTextMessage(ctx context.Context, em *pb.EmailMessage) (response *pb.SendResponse, err error) {
	response = &pb.SendResponse{Message: ""}
	err = m.sender.SendTextEmail(em)
	if err != nil {
		response.Message = err.Error()
	}

	return
}

func (m *mailServer) SendHTMLMessage(ctx context.Context, em *pb.EmailMessage) (response *pb.SendResponse, err error) {
	response = &pb.SendResponse{Message: ""}
	err = m.sender.SendHTMLEmail(em)
	if err != nil {
		response.Message = err.Error()
	}

	return
}

func main() {
	var sender Sender
	logger := logwrapper.New()

	if email, ok := os.LookupEnv("SMTP_EMAIL"); ok {
		sender = &SMTPSender{
			Email:    email,
			Password: os.Getenv("SMTP_PASSWORD"),
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Logger:   logger,
		}
	} else if key, ok := os.LookupEnv("SENDGRID_API_KEY"); ok {
		sender = &SendgridSender{Key: key, Logger: logger}
	} else {
		log.Fatalf("No available sender")
	}

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

	pb.RegisterTinyGomailServer(grpcServer, &mailServer{sender: sender})
	grpcServer.Serve(lis)
}
