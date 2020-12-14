package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/maxshend/tiny_gomail/tinygomail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const maxBufSize = 1048576

var lis *bufconn.Listener
var ctx = context.Background()

func init() {
	lis = bufconn.Listen(maxBufSize)
	s := grpc.NewServer()
	pb.RegisterTinyGomailServer(s, &mailServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func fakeConnection(t *testing.T) *grpc.ClientConn {
	t.Helper()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	return conn
}

func TestSendTextMessage(t *testing.T) {
	conn := fakeConnection(t)
	defer conn.Close()

	client := pb.NewTinyGomailClient(conn)
	_, err := client.SendTextMessage(ctx, &pb.EmailMessage{})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
}

func TestSendHTMLMessage(t *testing.T) {
	conn := fakeConnection(t)
	defer conn.Close()

	client := pb.NewTinyGomailClient(conn)
	_, err := client.SendHTMLMessage(ctx, &pb.EmailMessage{})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
}
