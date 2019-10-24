package main

import (
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {

	log.Printf("Doing Work")
	log.Printf("Sending Emails!")
	go sendMail("RECIPIENT")
	time.Sleep(time.Second)
	log.Printf("Done Sending Emails!")
}

func sendMail(recipient string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "EMAIL_SENDER")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>"+recipient+"</b> From Goroutine!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "GMAIL_USERNAME", "GMAIL_PASSWORD")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
