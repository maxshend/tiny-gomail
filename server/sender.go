package main

import pb "github.com/maxshend/tiny_gomail/tinygomail"

const htmlMIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"

// Sender describes an entity for sending emails
type Sender interface {
	SendTextEmail(em *pb.EmailMessage) (err error)
	SendHTMLEmail(em *pb.EmailMessage) (err error)
}
